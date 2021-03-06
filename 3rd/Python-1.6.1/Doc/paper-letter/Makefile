# you could set this to a4
PAPER=letter

# Where are the various programs?
LATEX=		latex
PDFLATEX=	pdflatex
DVIPS=		dvips -N0 -t $(PAPER)
MAKEINDEX=	makeindex -s ../texinputs/python.ist
PYTHON=		python

TOOLSDIR=	../tools
TEXINPUTS=	.:../texinputs:

MKDVI=		TEXINPUTS=$(TEXINPUTS) $(TOOLSDIR)/mkdvi.sh
# The --keep option is needed to keep the intermediate files for generating
# the HTML global module index.
MKHOWTO=	$(TOOLSDIR)/mkhowto --keep
MKPDF=		TEXINPUTS=$(TEXINPUTS) $(TOOLSDIR)/mkdvi.sh --pdf


# what's what
MANDVIFILES=	api.dvi ext.dvi lib.dvi ref.dvi tut.dvi
HOWTODVIFILES=	doc.dvi mac.dvi inst.dvi dist.dvi

MANPDFFILES=	api.pdf ext.pdf lib.pdf ref.pdf tut.pdf
HOWTOPDFFILES=	doc.pdf mac.pdf inst.pdf dist.pdf

MANPSFILES=	api.ps ext.ps lib.ps ref.ps tut.ps
HOWTOPSFILES=	doc.ps mac.ps inst.ps dist.ps

DVIFILES=	$(MANDVIFILES) $(HOWTODVIFILES)
PDFFILES=	$(MANPDFFILES) $(HOWTOPDFFILES)
PSFILES=	$(MANPSFILES) $(HOWTOPSFILES)


include ../Makefile.deps


SCRIPTS:	  ../tools/fix_hack ../tools/mkdvi.sh ../tools/indfix.py

$(DVIFILES):	  $(SCRIPTS)
$(PDFFILES):	  $(SCRIPTS)

$(HOWTODVIFILES): ../tools/mkhowto
$(HOWTOPDFFILES): ../tools/mkhowto


# Rules to build PostScript and PDF formats
.SUFFIXES: .dvi .ps .pdf

.dvi.ps:
	$(DVIPS) -o $@ $<


all:	ps
dvi:	$(DVIFILES)
pdf:	$(PDFFILES)
ps:	$(PSFILES)


README:	$(PSFILES) $(TOOLSDIR)/getpagecounts
	$(TOOLSDIR)/getpagecounts >$@


# Python/C API Reference Manual
api.dvi: api.tex $(APIFILES)
	$(MKDVI) api

api.pdf: api.tex $(APIFILES)
	$(MKPDF) api

api.tex: ../api/api.tex ../api/refcounts.dat ../tools/anno-api.py
	$(PYTHON) ../tools/anno-api.py -o $@ ../api/api.tex

# Distributing Python Modules
dist.dvi: $(DISTFILES)
	$(MKHOWTO) --dvi ../dist/dist.tex

dist.pdf: $(DISTFILES)
	$(MKHOWTO) --pdf ../dist/dist.tex

# Documenting Python
doc.dvi: $(DOCFILES)
	$(MKHOWTO) --dvi ../doc/doc.tex

doc.pdf: $(DOCFILES)
	$(MKHOWTO) --pdf ../doc/doc.tex

# Extending and Embedding the Python Interpreter
ext.dvi: $(EXTFILES)
	$(MKDVI) ext

ext.pdf: $(EXTFILES)
	$(MKPDF) ext

# Installing Python Modules
inst.dvi: $(INSTFILES)
	$(MKHOWTO) --dvi ../inst/inst.tex

inst.pdf: $(INSTFILES)
	$(MKHOWTO) --pdf ../inst/inst.tex

# Python Library Reference
lib.dvi: $(LIBFILES)
	$(MKDVI) lib

lib.pdf: $(LIBFILES)
	$(MKPDF) lib

# Macintosh Library Modules
mac.dvi: $(MACFILES)
	$(MKHOWTO) --dvi ../mac/mac.tex

mac.pdf: $(MACFILES)
	$(MKHOWTO) --pdf ../mac/mac.tex

# Python Reference Manual
ref.dvi: $(REFFILES)
	$(MKDVI) ref

ref.pdf: $(REFFILES)
	$(MKPDF) ref

# Python Tutorial
tut.dvi: $(TUTFILES)
	$(MKDVI) tut

tut.pdf: $(TUTFILES)
	$(MKPDF) tut


clean:
	rm -f *~ *.aux *.idx *.ilg *.ind *.log *.toc *.bkm *.syn *.pla api.tex

clobber: clean
	rm -f $(DVIFILES) $(PSFILES) $(PDFFILES)

realclean: clobber
distclean: clobber
