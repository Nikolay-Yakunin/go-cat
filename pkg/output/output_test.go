package output

import (
    "os"
    "testing"

    "github.com/Nikolay-Yakunin/go-cat/pkg/flags"
)

func TestProcessFile(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        flags    flags.Flags
    }{
        {
            name:   "Test with valid input and no flags",
            input:  "line1\nline2\n",
            flags:  flags.Flags{},
        },
        {
            name:   "Test with -n flag",
            input:  "line1\nline2\n",
            flags:  flags.Flags{FlagN: true},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Создаем временный файл
            tmpFile, err := os.CreateTemp("", "testfile")
            if err != nil {
                t.Fatalf("failed to create temp file: %v", err)
            }
            defer os.Remove(tmpFile.Name()) // Удаляем файл после теста
            defer tmpFile.Close()           // Закрываем файл после использования

            // Записываем входные данные в временный файл
            _, err = tmpFile.WriteString(tt.input)
            if err != nil {
                t.Fatalf("failed to write to temp file: %v", err)
            }

            // Возвращаем указатель на файл в начало для чтения
            _, err = tmpFile.Seek(0, 0)
            if err != nil {
                t.Fatalf("failed to seek to the beginning of the file: %v", err)
            }

            // Вызываем ProcessFile с временным файлом
            ProcessFile(tmpFile, tt.flags)

            // Вывод будет отправлен в os.Stdout (консоль)
        })
    }
}

func TestPrintLine(t *testing.T) {
    tests := []struct {
        name        string
        rawLine     string
        flags       flags.Flags
        lineNum     int
        prevBlank   bool
        newPrevBlank bool
    }{
        {
            name:        "Test print valid line with -n flag",
            rawLine:     "line1",
            flags:       flags.Flags{FlagN: true},
            lineNum:     1,
            prevBlank:   false,
            newPrevBlank: false,
        },
        {
            name:        "Test print empty line with -b flag",
            rawLine:     "",
            flags:       flags.Flags{FlagB: true},
            lineNum:     1,
            prevBlank:   false,
            newPrevBlank: true,
        },
        {
            name:        "Test print non-empty line with -b flag",
            rawLine:     "line1",
            flags:       flags.Flags{FlagB: true},
            lineNum:     1,
            prevBlank:   false,
            newPrevBlank: false,
        },
        {
            name:        "Test print line with -t and -e flags",
            rawLine:     "line1\t",
            flags:       flags.Flags{FlagT: true, FlagE: true},
            lineNum:     1,
            prevBlank:   false,
            newPrevBlank: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Вызываем printLine напрямую с os.Stdout

            printLine(tt.rawLine, tt.flags, &tt.lineNum, &tt.prevBlank)

            // Проверяем изменение previousBlank
            if tt.prevBlank != tt.newPrevBlank {
                t.Errorf("expected previousBlank to be %v, got %v", tt.newPrevBlank, tt.prevBlank)
            }
        })
    }
}