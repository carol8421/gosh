// mkerrors.sh
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

// godefs -gsyscall _errors.c

// MACHINE GENERATED - DO NOT EDIT.

package syscall

// Constants
const (
	EMULTIHOP	= 0x48;
	EUNATCH		= 0x31;
	EAFNOSUPPORT	= 0x61;
	EREMCHG		= 0x4e;
	EACCES		= 0xd;
	EL3RST		= 0x2f;
	EDESTADDRREQ	= 0x59;
	EILSEQ		= 0x54;
	ESPIPE		= 0x1d;
	EMLINK		= 0x1f;
	EOWNERDEAD	= 0x82;
	ENOTTY		= 0x19;
	EBADE		= 0x34;
	EBADF		= 0x9;
	EBADR		= 0x35;
	EADV		= 0x44;
	ERANGE		= 0x22;
	ECANCELED	= 0x7d;
	ETXTBSY		= 0x1a;
	ENOMEM		= 0xc;
	EINPROGRESS	= 0x73;
	ENOTEMPTY	= 0x27;
	ENOTBLK		= 0xf;
	EPROTOTYPE	= 0x5b;
	ERESTART	= 0x55;
	EISNAM		= 0x78;
	ENOMSG		= 0x2a;
	EALREADY	= 0x72;
	ETIMEDOUT	= 0x6e;
	ENODATA		= 0x3d;
	EINTR		= 0x4;
	ENOLINK		= 0x43;
	EPERM		= 0x1;
	ELOOP		= 0x28;
	ENETDOWN	= 0x64;
	ESTALE		= 0x74;
	ENOTSOCK	= 0x58;
	ENOSR		= 0x3f;
	ECHILD		= 0xa;
	ELNRNG		= 0x30;
	EPIPE		= 0x20;
	EBADMSG		= 0x4a;
	EBFONT		= 0x3b;
	EREMOTE		= 0x42;
	ETOOMANYREFS	= 0x6d;
	EPFNOSUPPORT	= 0x60;
	ENONET		= 0x40;
	EXFULL		= 0x36;
	EBADSLT		= 0x39;
	ENOTNAM		= 0x76;
	ENOCSI		= 0x32;
	EADDRINUSE	= 0x62;
	ENETRESET	= 0x66;
	EISDIR		= 0x15;
	EIDRM		= 0x2b;
	ECOMM		= 0x46;
	EBADFD		= 0x4d;
	EL2HLT		= 0x33;
	ENOKEY		= 0x7e;
	EINVAL		= 0x16;
	ESHUTDOWN	= 0x6c;
	EKEYREJECTED	= 0x81;
	ELIBSCN		= 0x51;
	ENAVAIL		= 0x77;
	EOVERFLOW	= 0x4b;
	EUCLEAN		= 0x75;
	ENOMEDIUM	= 0x7b;
	EBUSY		= 0x10;
	EPROTO		= 0x47;
	ENODEV		= 0x13;
	EKEYEXPIRED	= 0x7f;
	EROFS		= 0x1e;
	ELIBACC		= 0x4f;
	E2BIG		= 0x7;
	EDEADLK		= 0x23;
	ENOTDIR		= 0x14;
	ECONNRESET	= 0x68;
	ENXIO		= 0x6;
	EBADRQC		= 0x38;
	ENAMETOOLONG	= 0x24;
	ESOCKTNOSUPPORT	= 0x5e;
	ELIBEXEC	= 0x53;
	EDOTDOT		= 0x49;
	EADDRNOTAVAIL	= 0x63;
	ETIME		= 0x3e;
	EPROTONOSUPPORT	= 0x5d;
	ENOTRECOVERABLE	= 0x83;
	EIO		= 0x5;
	ENETUNREACH	= 0x65;
	EXDEV		= 0x12;
	EDQUOT		= 0x7a;
	EREMOTEIO	= 0x79;
	ENOSPC		= 0x1c;
	ENOEXEC		= 0x8;
	EMSGSIZE	= 0x5a;
	EDOM		= 0x21;
	ENOSTR		= 0x3c;
	EFBIG		= 0x1b;
	ESRCH		= 0x3;
	ECHRNG		= 0x2c;
	EHOSTDOWN	= 0x70;
	ENOLCK		= 0x25;
	ENFILE		= 0x17;
	ENOSYS		= 0x26;
	ENOTCONN	= 0x6b;
	ENOTSUP		= 0x5f;
	ESRMNT		= 0x45;
	EDEADLOCK	= 0x23;
	ECONNABORTED	= 0x67;
	ENOANO		= 0x37;
	EISCONN		= 0x6a;
	EUSERS		= 0x57;
	ENOPROTOOPT	= 0x5c;
	EMFILE		= 0x18;
	ENOBUFS		= 0x69;
	EL3HLT		= 0x2e;
	EFAULT		= 0xe;
	EWOULDBLOCK	= 0xb;
	ELIBBAD		= 0x50;
	ESTRPIPE	= 0x56;
	ECONNREFUSED	= 0x6f;
	EAGAIN		= 0xb;
	ELIBMAX		= 0x52;
	EEXIST		= 0x11;
	EL2NSYNC	= 0x2d;
	ENOENT		= 0x2;
	ENOPKG		= 0x41;
	EKEYREVOKED	= 0x80;
	EHOSTUNREACH	= 0x71;
	ENOTUNIQ	= 0x4c;
	EOPNOTSUPP	= 0x5f;
	EMEDIUMTYPE	= 0x7c;
	SIGBUS		= 0x7;
	SIGTTIN		= 0x15;
	SIGPROF		= 0x1b;
	SIGFPE		= 0x8;
	SIGHUP		= 0x1;
	SIGTTOU		= 0x16;
	SIGSTKFLT	= 0x10;
	SIGUSR1		= 0xa;
	SIGURG		= 0x17;
	SIGIO		= 0x1d;
	SIGQUIT		= 0x3;
	SIGCLD		= 0x11;
	SIGABRT		= 0x6;
	SIGTRAP		= 0x5;
	SIGVTALRM	= 0x1a;
	SIGPOLL		= 0x1d;
	SIGSEGV		= 0xb;
	SIGCONT		= 0x12;
	SIGPIPE		= 0xd;
	SIGWINCH	= 0x1c;
	SIGXFSZ		= 0x19;
	SIGCHLD		= 0x11;
	SIGSYS		= 0x1f;
	SIGSTOP		= 0x13;
	SIGALRM		= 0xe;
	SIGUSR2		= 0xc;
	SIGTSTP		= 0x14;
	SIGKILL		= 0x9;
	SIGXCPU		= 0x18;
	SIGUNUSED	= 0x1f;
	SIGPWR		= 0x1e;
	SIGILL		= 0x4;
	SIGINT		= 0x2;
	SIGIOT		= 0x6;
	SIGTERM		= 0xf;
	O_EXCL		= 0x80;
)

// Types


// Error table
var errors = [...]string{
	72: "multihop attempted",
	49: "protocol driver not attached",
	97: "address family not supported by protocol",
	78: "remote address changed",
	13: "permission denied",
	47: "level 3 reset",
	89: "destination address required",
	84: "invalid or incomplete multibyte or wide character",
	29: "illegal seek",
	31: "too many links",
	130: "owner died",
	25: "inappropriate ioctl for device",
	52: "invalid exchange",
	9: "bad file descriptor",
	53: "invalid request descriptor",
	68: "advertise error",
	34: "numerical result out of range",
	125: "operation canceled",
	26: "text file busy",
	12: "cannot allocate memory",
	115: "operation now in progress",
	39: "directory not empty",
	15: "block device required",
	91: "protocol wrong type for socket",
	85: "interrupted system call should be restarted",
	120: "is a named type file",
	42: "no message of desired type",
	114: "operation already in progress",
	110: "connection timed out",
	61: "no data available",
	4: "interrupted system call",
	67: "link has been severed",
	1: "operation not permitted",
	40: "too many levels of symbolic links",
	100: "network is down",
	116: "stale NFS file handle",
	88: "socket operation on non-socket",
	63: "out of streams resources",
	10: "no child processes",
	48: "link number out of range",
	32: "broken pipe",
	74: "bad message",
	59: "bad font file format",
	66: "object is remote",
	109: "too many references: cannot splice",
	96: "protocol family not supported",
	64: "machine is not on the network",
	54: "exchange full",
	57: "invalid slot",
	118: "not a XENIX named type file",
	50: "no CSI structure available",
	98: "address already in use",
	102: "network dropped connection on reset",
	21: "is a directory",
	43: "identifier removed",
	70: "communication error on send",
	77: "file descriptor in bad state",
	51: "level 2 halted",
	126: "required key not available",
	22: "invalid argument",
	108: "cannot send after transport endpoint shutdown",
	129: "key was rejected by service",
	81: ".lib section in a.out corrupted",
	119: "no XENIX semaphores available",
	75: "value too large for defined data type",
	117: "structure needs cleaning",
	123: "no medium found",
	16: "device or resource busy",
	71: "protocol error",
	19: "no such device",
	127: "key has expired",
	30: "read-only file system",
	79: "can not access a needed shared library",
	7: "argument list too long",
	35: "resource deadlock avoided",
	20: "not a directory",
	104: "connection reset by peer",
	6: "no such device or address",
	56: "invalid request code",
	36: "file name too long",
	94: "socket type not supported",
	83: "cannot exec a shared library directly",
	73: "RFS specific error",
	99: "cannot assign requested address",
	62: "timer expired",
	93: "protocol not supported",
	131: "state not recoverable",
	5: "input/output error",
	101: "network is unreachable",
	18: "invalid cross-device link",
	122: "disk quota exceeded",
	121: "remote I/O error",
	28: "no space left on device",
	8: "exec format error",
	90: "message too long",
	33: "numerical argument out of domain",
	60: "device not a stream",
	27: "file too large",
	3: "no such process",
	44: "channel number out of range",
	112: "host is down",
	37: "no locks available",
	23: "too many open files in system",
	38: "function not implemented",
	107: "transport endpoint is not connected",
	95: "operation not supported",
	69: "srmount error",
	103: "software caused connection abort",
	55: "no anode",
	106: "transport endpoint is already connected",
	87: "too many users",
	92: "protocol not available",
	24: "too many open files",
	105: "no buffer space available",
	46: "level 3 halted",
	14: "bad address",
	11: "resource temporarily unavailable",
	80: "accessing a corrupted shared library",
	86: "streams pipe error",
	111: "connection refused",
	82: "attempting to link in too many shared libraries",
	17: "file exists",
	45: "level 2 not synchronized",
	2: "no such file or directory",
	65: "package not installed",
	128: "key has been revoked",
	113: "no route to host",
	76: "name not unique on network",
	124: "wrong medium type",
}
