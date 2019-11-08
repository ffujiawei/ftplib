package ftplib

// Commands, defined in RFC 959
const (
	ABOR = "ABOR" // Abort an active file transfer.
	ACCT = "ACCT" // Account information.
	ADAT = "ADAT" // Authentication/Security Data
	ALLO = "ALLO" // Allocate sufficient disk space to receive a file.
	APPE = "APPE" // Append (with create)
	AUTH = "AUTH" // Authentication/Security Mechanism
	AVBL = "AVBL" // Get the available space
	CCC  = "CCC"  // Clear Command Channel
	CDUP = "CDUP" // Change to Parent Directory.
	CONF = "CONF" // Confidentiality Protection Command
	CSID = "CSID" // Client / Server Identification
	CWD  = "CWD"  // Change working directory.
	DELE = "DELE" // Delete file.
	DSIZ = "DSIZ" // Get the directory size
	ENC  = "ENC"  // Privacy Protected Channel
	EPRT = "EPRT" // Specifies an extended address and port to which the server should connect.
	EPSV = "EPSV" // Enter extended passive mode.
	FEAT = "FEAT" // Get the feature list implemented by the server.
	HELP = "HELP" // Returns usage documentation on a command if specified, else a general help document is returned.
	HOST = "HOST" // Identify desired virtual host on server, by name.
	LANG = "LANG" // Language Negotiation
	LIST = "LIST" // Returns information of a file or directory if specified, else information of the current working directory is returned.
	LPRT = "LPRT" // Specifies a long address and port to which the server should connect.
	LPSV = "LPSV" // Enter long passive mode.
	MDTM = "MDTM" // Return the last-modified time of a specified file.
	MFCT = "MFCT" // Modify the creation time of a file.
	MFF  = "MFF"  // Modify fact (the last modification time, creation time, UNIX group/owner/mode of a file).
	MFMT = "MFMT" // Modify the last modification time of a file.
	MIC  = "MIC"  // Integrity Protected Command
	MKD  = "MKD"  // Make directory.
	MLSD = "MLSD" // Lists the contents of a directory if a directory is named.
	MLST = "MLST" // Provides data about exactly the object named on its command line, and no others.
	MODE = "MODE" // Sets the transfer mode (Stream, Block, or Compressed).
	NLST = "NLST" // Returns a list of file names in a specified directory.
	NOOP = "NOOP" // No operation (dummy packet; used mostly on keepalives).
	OPTS = "OPTS" // Select options for a feature (for example OPTS UTF8 ON).
	PASS = "PASS" // Authentication password.
	PASV = "PASV" // Enter passive mode.
	PBSZ = "PBSZ" // Protection Buffer Size
	PORT = "PORT" // Specifies an address and port to which the server should connect.
	PROT = "PROT" // Data Channel Protection Level.
	PWD  = "PWD"  // Print working directory. Returns the current directory of the host.
	QUIT = "QUIT" // Disconnect.
	REIN = "REIN" // Re initializes the connection.
	REST = "REST" // Restart transfer from the specified point.
	RETR = "RETR" // Retrieve a copy of the file
	RMD  = "RMD"  // Remove a directory.
	RMDA = "RMDA" // Remove a directory tree
	RNFR = "RNFR" // Rename from.
	RNTO = "RNTO" // Rename to.
	SITE = "SITE" // Sends site specific commands to remote server (like SITE IDLE 60or SITE UMASK 002). Inspect SITE HELP output for complete list of supported commands.
	SIZE = "SIZE" // Return the size of a file.
	SMNT = "SMNT" // Mount file structure.
	SPSV = "SPSV" // Use single port passive mode (only one TCP port number for both control connections and passive-mode data connections)
	STAT = "STAT" // Returns the current status.
	STOR = "STOR" // Accept the data and to store the data as a file at the server site
	STOU = "STOU" // Store file uniquely.
	STRU = "STRU" // Set file transfer structure.
	SYST = "SYST" // Return system type.
	THMB = "THMB" // Get a thumbnail of a remote image file
	TYPE = "TYPE" // Sets the transfer mode (ASCII/Binary).
	USER = "USER" // Authentication username.
	XCUP = "XCUP" // Change to the parent of the current working directory
	XMKD = "XMKD" // Make a directory
	XPWD = "XPWD" // Print the current working directory
	XRCP = "XRCP" //
	XRMD = "XRMD" // Remove the directory
	XRSQ = "XRSQ" //
	XSEM = "XSEM" // Send, mail if cannot
	XSEN = "XSEN" // Send to terminal
)
