package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Структура для хранения конфигурации программы
type config struct {
	fields         []int
	delimiter      string
	hideWrongLines bool
}

// Функция для парсинга флагов
func parseFlags() config {
	cfg := config{}
	fieldsFlag := flag.String("f", "", "выбрать поля (колонки), перечисленные через запятую")
	flag.StringVar(&cfg.delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&cfg.hideWrongLines, "s", false, "только строки с разделителем")
	flag.Parse()

	// Разбор поля fields
	if *fieldsFlag != "" {
		for _, f := range strings.Split(*fieldsFlag, ",") {
			var field int
			_, err := fmt.Sscanf(f, "%d", &field)
			if err != nil {
				fmt.Fprintf(os.Stderr, "неправильный формат поля: %v\n", err)
				os.Exit(1)
			}
			// Уменьшаем поле на 1 для приведения к индексу
			cfg.fields = append(cfg.fields, field-1)
		}
	}

	return cfg
}

// Функция для обработки строк из stdin и вывода выбранных колонок
func cut(sc *bufio.Scanner, cfg config) {
	for sc.Scan() {
		line := sc.Text()

		// Пропускаем строки без разделителя, если установлен флаг -s
		if cfg.hideWrongLines && !strings.Contains(line, cfg.delimiter) {
			continue
		}

		// Разделяем строку на колонки
		columns := strings.Split(line, cfg.delimiter)

		// Если поля не указаны, выводим все колонки
		if len(cfg.fields) == 0 {
			fmt.Println(strings.Join(columns, " "))
		} else {
			// Выбираем только указанные колонки
			var selectedFields []string
			for _, field := range cfg.fields {
				if field >= 0 && field < len(columns) {
					selectedFields = append(selectedFields, columns[field])
				}
			}
			// Выводим выбранные колонки, если они есть
			if len(selectedFields) > 0 {
				fmt.Println(strings.Join(selectedFields, cfg.delimiter))
			}
		}
	}
}

func main() {
	// Парсинг флагов
	cfg := parseFlags()

	// Создаем сканер для чтения из stdin
	sc := bufio.NewScanner(os.Stdin)

	// Обрабатываем входные строки с использованием конфигурации
	cut(sc, cfg)
}
