package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init() {
	// Fungsi logger disini akan diinisialisasi saat package ini di-load
	// Contoh: setup konfigurasi logger, inisialisasi file log, dsb
	// Misalnya, kita bisa setup logger untuk mencetak ke konsol atau file
	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.DebugLevel)
}