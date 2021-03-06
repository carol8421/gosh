#
# Secret Labs' Regular Expression Engine
#
# re-compatible interface for the sre matching engine
#
# Copyright (c) 1998-2000 by Secret Labs AB.  All rights reserved.
#
# This version of the SRE library can be redistributed under CNRI's
# Python 1.6 license.  For any other use, please contact Secret Labs
# AB (info@pythonware.com).
#
# Portions of this engine have been developed in cooperation with
# CNRI.  Hewlett-Packard provided funding for 1.6 integration and
# other compatibility work.
#

# FIXME: change all FIXME's to XXX ;-)

import sre_compile
import sre_parse

import string

# flags
I = IGNORECASE = sre_compile.SRE_FLAG_IGNORECASE
L = LOCALE = sre_compile.SRE_FLAG_LOCALE
M = MULTILINE = sre_compile.SRE_FLAG_MULTILINE
S = DOTALL = sre_compile.SRE_FLAG_DOTALL
X = VERBOSE = sre_compile.SRE_FLAG_VERBOSE

# sre extensions (may or may not be in 1.6/2.0 final)
T = TEMPLATE = sre_compile.SRE_FLAG_TEMPLATE
U = UNICODE = sre_compile.SRE_FLAG_UNICODE

# sre exception
error = sre_compile.error

# --------------------------------------------------------------------
# public interface

# FIXME: add docstrings

def match(pattern, string, flags=0):
    return _compile(pattern, flags).match(string)

def search(pattern, string, flags=0):
    return _compile(pattern, flags).search(string)

def sub(pattern, repl, string, count=0):
    return _compile(pattern, 0).sub(repl, string, count)

def subn(pattern, repl, string, count=0):
    return _compile(pattern, 0).subn(repl, string, count)

def split(pattern, string, maxsplit=0):
    return _compile(pattern, 0).split(string, maxsplit)

def findall(pattern, string, maxsplit=0):
    return _compile(pattern, 0).findall(string, maxsplit)

def compile(pattern, flags=0):
    return _compile(pattern, flags)

def purge():
    _cache.clear()

def template(pattern, flags=0):
    return _compile(pattern, flags|T)

def escape(pattern):
    s = list(pattern)
    for i in range(len(pattern)):
        c = pattern[i]
        if not ("a" <= c <= "z" or "A" <= c <= "Z" or "0" <= c <= "9"):
            if c == "\000":
                s[i] = "\\000"
            else:
                s[i] = "\\" + c
    return _join(s, pattern)

# --------------------------------------------------------------------
# internals

_cache = {}
_MAXCACHE = 100

def _join(seq, sep):
    # internal: join into string having the same type as sep
    return string.join(seq, sep[:0])

def _compile(*key):
    # internal: compile pattern
    p = _cache.get(key)
    if p is not None:
        return p
    pattern, flags = key
    if type(pattern) not in sre_compile.STRING_TYPES:
        return pattern
    try:
        p = sre_compile.compile(pattern, flags)
    except error, v:
        raise error, v # invalid expression
    if len(_cache) >= _MAXCACHE:
        _cache.clear()
    _cache[key] = p
    return p

def _sub(pattern, template, string, count=0):
    # internal: pattern.sub implementation hook
    return _subn(pattern, template, string, count)[0]

def _subn(pattern, template, string, count=0):
    # internal: pattern.subn implementation hook
    if callable(template):
        filter = template
    else:
        template = sre_parse.parse_template(template, pattern)
        def filter(match, template=template):
            return sre_parse.expand_template(template, match)
    n = i = 0
    s = []
    append = s.append
    c = pattern.scanner(string)
    while not count or n < count:
        m = c.search()
        if not m:
            break
        b, e = m.span()
        if i < b:
            append(string[i:b])
        append(filter(m))
        i = e
        n = n + 1
    append(string[i:])
    return _join(s, string[:0]), n

def _split(pattern, string, maxsplit=0):
    # internal: pattern.split implementation hook
    n = i = 0
    s = []
    append = s.append
    extend = s.extend
    c = pattern.scanner(string)
    g = pattern.groups
    while not maxsplit or n < maxsplit:
        m = c.search()
        if not m:
            break
        b, e = m.span()
        if b == e:
            if i >= len(string):
                break
            continue
        append(string[i:b])
        if g and b != e:
            extend(m.groups())
        i = e
        n = n + 1
    append(string[i:])
    return s

# register myself for pickling

import copy_reg

def _pickle(p):
    return _compile, (p.pattern, p.flags)

copy_reg.pickle(type(_compile("", 0)), _pickle, _compile)

# --------------------------------------------------------------------
# experimental stuff (see python-dev discussions for details)

class Scanner:
    def __init__(self, lexicon):
        from sre_constants import BRANCH, SUBPATTERN
        self.lexicon = lexicon
        # combine phrases into a compound pattern
        p = []
        s = sre_parse.Pattern()
        for phrase, action in lexicon:
            p.append(sre_parse.SubPattern(s, [
                (SUBPATTERN, (len(p), sre_parse.parse(phrase))),
                ]))
        p = sre_parse.SubPattern(s, [(BRANCH, (None, p))])
        s.groups = len(p)
        self.scanner = sre_compile.compile(p)
    def scan(self, string):
        result = []
        append = result.append
        match = self.scanner.match
        i = 0
        while 1:
            m = match(string, i)
            if not m:
                break
            j = m.end()
            if i == j:
                break
            action = self.lexicon[m.lastindex][1]
            if callable(action):
                self.match = match
                action = action(self, m.group())
            if action is not None:
                append(action)
            i = j
        return result, string[i:]
