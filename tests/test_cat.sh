# Путь к бинарнику go-cat
GO_CAT_BIN="../build/linux/go-cat"

# Файлы для тестов
TEST_FILE="test_file.txt"
GO_CAT_OUTPUT="go-cat_output.txt" 
CAT_OUTPUT="cat_output.txt"       

echo "Hello\n\nWorld!\n\tTabbed line\n\n\nNew line\n" > "$TEST_FILE"

run_test() {
    description="$1"
    flags="$2"
    echo "Running test: $description with flags: $flags"

    "$GO_CAT_BIN" $flags "$TEST_FILE" > "$GO_CAT_OUTPUT"

    cat $flags "$TEST_FILE" > "$CAT_OUTPUT"

    diff_output=$(diff -u "$CAT_OUTPUT" "$GO_CAT_OUTPUT")
    if [ -z "$diff_output" ]; then
        echo "Test passed!"
    else
        echo "Test failed! Differences:"
        echo "$diff_output"
    fi
    echo
}

# Тесты для отдельных флагов
run_test "Flag -b (number non-empty lines)" "-b"
run_test "Flag -e (show \$ at end of each line)" "-e"
run_test "Flag -n (number all lines)" "-n"
run_test "Flag -s (squeeze blank lines)" "-s"
run_test "Flag -t (show tabs as ^I)" "-t"
run_test "Flag -v (show non-printable characters)" "-v"

# Тесты для комбинаций флагов
run_test "Flags -s and -n (squeeze blank lines, number all lines)" "-s -n"
run_test "Flags -b and -e (number non-empty lines, show \$ at end)" "-b -e"
run_test "Flags -t and -e (show tabs as ^I, show \$ at end)" "-t -e"
run_test "Flags -b, -s, and -n (number non-empty lines, squeeze blank lines, number all lines)" "-b -s -n"

# Тесты для GNU-стиля флагов
run_test "GNU Flag --number-nonblank (equivalent to -b)" "--number-nonblank"
run_test "GNU Flag --squeeze-blank (equivalent to -s)" "--squeeze-blank"
run_test "GNU Flag --number (equivalent to -n)" "--number"
run_test "GNU Flag -E (show \$ at end of each line, without -v)" "-E"
run_test "GNU Flag -T (show tabs as ^I, without -v)" "-T"

# Тесты GNU с короткими флагами
run_test "Flags -n and --squeeze-blank" "-n --squeeze-blank"
run_test "Flags -b and -E" "-b -E"
run_test "Flags --number-nonblank and -T" "--number-nonblank -T"
run_test "Flags --squeeze-blank and -e" "--squeeze-blank -e"
run_test "Flags -T and -e" "-T -e"
run_test "Flags --number-nonblank, -s, and -n" "--number-nonblank -s -n"

rm -f "$GO_CAT_OUTPUT" "$CAT_OUTPUT" "$TEST_FILE"