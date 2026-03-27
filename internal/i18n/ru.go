package i18n

var ruTranslations = map[string]string{
	// Navigation
	"Summary":      "Обзор",
	"Databases":    "Базы данных",
	"Destinations": "Хранилища",
	"Backup tasks": "Задачи резервного копирования",
	"Executions":   "Выполнения",
	"Restorations": "Восстановления",
	"Webhooks":     "Вебхуки",
	"Profile":      "Профиль",
	"About":        "О программе",

	// Header/Auth
	"Log out":  "Выйти",
	"Theme":    "Тема",
	"System":   "Системная",
	"Light":    "Светлая",
	"Dark":     "Тёмная",
	"Language": "Язык",
	"English":  "English",
	"Russian":  "Русский",

	// Auth pages
	"Create first user":        "Создать первого пользователя",
	"Full name":                "Полное имя",
	"Email":                    "Email",
	"Password":                 "Пароль",
	"Confirm password":         "Подтвердите пароль",
	"Create user and continue": "Создать пользователя и продолжить",
	"Login":                    "Вход",
	"Remember me":              "Запомнить меня",

	// Buttons
	"Previous":                 "Назад",
	"Next":                     "Далее",
	"Save":                     "Сохранить",
	"Cancel":                   "Отмена",
	"Delete":                   "Удалить",
	"Edit":                     "Редактировать",
	"Back":                     "Назад",
	"Submit":                   "Отправить",
	"Close":                    "Закрыть",
	"Run backup now":           "Запустить бэкап сейчас",
	"Test connection":          "Проверить подключение",
	"Show executions":          "Показать выполнения",
	"Add database":             "Добавить базу данных",
	"Add destination":          "Добавить хранилище",
	"Create backup":            "Создать задачу бэкапа",
	"Duplicate backup task":    "Дублировать задачу",
	"Delete backup task":       "Удалить задачу",
	"Edit backup task":         "Редактировать задачу",
	"Delete database":          "Удалить базу данных",
	"Edit database":            "Редактировать базу данных",
	"Delete destination":       "Удалить хранилище",
	"Edit destination":         "Редактировать хранилище",
	"Restore backup":           "Восстановить из бэкапа",
	"Connect PostgreSQL":       "Подключить PostgreSQL",
	"Connect and import":       "Подключить и импортировать",
	"Destinations description": "Здесь вы можете управлять S3 хранилищами. Вы можете пропустить создание S3 хранилища, если хотите использовать локальное хранилище для бэкапов.",
	"Health check description": "Проверка работоспособности баз данных и хранилищ выполняется автоматически каждые 10 минут, при запуске PG Back Web и при нажатии кнопки \"Проверить подключение\" для каждого ресурса.",

	// Status
	"Status":    "Статус",
	"Active":    "Активно",
	"Inactive":  "Неактивно",
	"Running":   "Выполняется",
	"Success":   "Успешно",
	"Failed":    "Ошибка",
	"Deleted":   "Удалено",
	"Healthy":   "Работает",
	"Unhealthy": "Проблемы",
	"Pending":   "Ожидание",
	"Completed": "Завершено",
	"Warning":   "Предупреждение",
	"Error":     "Ошибка",
	"Info":      "Информация",
	"Yes":       "Да",
	"No":        "Нет",

	// Table headers
	"Name":              "Название",
	"Database":          "База данных",
	"Destination":       "Хранилище",
	"Backup":            "Бэкап",
	"Schedule":          "Расписание",
	"Retention":         "Хранение",
	"Created at":        "Создано",
	"Updated at":        "Обновлено",
	"Started at":        "Начато",
	"Finished at":       "Завершено",
	"Duration":          "Длительность",
	"File size":         "Размер файла",
	"Version":           "Версия",
	"PostgreSQL":        "PostgreSQL",
	"Connection string": "Строка подключения",
	"Bucket name":       "Имя бакета",
	"Endpoint":          "Эндпоинт",
	"Region":            "Регион",
	"Access key":        "Ключ доступа",
	"Secret key":        "Секретный ключ",
	"Resource":          "Ресурс",
	"Total":             "Всего",
	"Message":           "Сообщение",
	"Took":              "Затрачено",
	"ID":                "ID",
	"Execution":         "Выполнение",

	// Form labels and placeholders
	"Description":          "Описание",
	"Search":               "Поиск",
	"Filter":               "Фильтр",
	"Actions":              "Действия",
	"Details":              "Подробности",
	"Options":              "Опции",
	"Learn more":           "Узнать больше",
	"Select an event type": "Выберите тип события",

	// Health
	"Health status": "Состояние системы",

	// Backup options
	"--data-only":   "Только данные",
	"--schema-only": "Только схема",
	"--clean":       "Очистка",
	"--if-exists":   "Если существует",
	"--create":      "Создать",
	"--no-comments": "Без комментариев",

	// Misc
	"Chart waiting for data": "График ожидает данные",
	"Quantity":               "Количество",
	"Days":                   "дней",

	// Restore
	"Existing database": "Существующая база данных",
	"Other database":    "Другая база данных",

	// Webhooks
	"POST": "POST",
	"GET":  "GET",

	// Profile
	"Update profile":       "Обновить профиль",
	"Current password":     "Текущий пароль",
	"New password":         "Новый пароль",
	"Confirm new password": "Подтвердите новый пароль",

	// About page
	"License":          "Лицензия",
	"Repository":       "Репозиторий",
	"About the author": "Об авторе",

	// Destinations page
	"S3 Destinations": "S3 хранилища",

	// Common
	"Loading":       "Загрузка...",
	"No results":    "Нет результатов",
	"Confirm":       "Подтвердить",
	"Are you sure?": "Вы уверены?",
	"Are you sure you want to delete this database?": "Вы уверены, что хотите удалить эту базу данных?",

	// Local backups section
	"Local backups":                 "Локальные бэкапы",
	"Remote backups":                "Удалённые бэкапы",
	"Examples & common expressions": "Примеры и распространённые выражения",

	// Timezone
	"Timezone": "Часовой пояс",

	// Health
	"Health check info": "Проверка выполняется каждые 10 минут",

	// Webhooks
	"No executions found":                         "Выполнений не найдено",
	"Wait for the first execution to appear here": "Подождите первого выполнения",

	// Chart
	"No data": "Нет данных",

	"Local backups are stored on the server where PG Back Web is running. They are stored under the /backups directory, so you can mount a Docker volume to this directory to persist backups in any way you want.":                                                       "Локальные бэкапы хранятся на сервере, где запущен PG Back Web. Они сохраняются в каталоге /backups, поэтому вы можете примонтировать Docker-том к этому каталогу и хранить бэкапы удобным для вас способом.",
	"Remote backups are stored in a destination. A destination is an S3-compatible remote storage. With this option, you do not need to worry about creating and managing Docker volumes.":                                                                                "Удалённые бэкапы хранятся в хранилище назначения. Назначение — это удалённое S3-совместимое хранилище. С этим вариантом не нужно создавать и управлять Docker-томами.",
	"A cron expression is a string used to define a schedule for running tasks in Unix-like operating systems. It consists of five fields representing minute, hour, day of month, month, and day of week. Cron expressions enable precise scheduling of periodic tasks.": "Cron-выражение — это строка для задания расписания запуска задач в Unix-подобных системах. Оно состоит из пяти полей: минута, час, день месяца, месяц и день недели. Cron-выражения позволяют точно планировать периодические задачи.",
	"This is the time zone in which the cron expression will be evaluated.": "Это часовой пояс, в котором будет вычисляться cron-выражение.",
	"Backup filenames will always use the server timezone (currently %s).":  "Имена файлов бэкапов всегда используют часовой пояс сервера (сейчас %s).",
	"The destination directory is where backups will be stored. This directory is relative to the destination base directory. It should start with a slash, contain no spaces, and should not end with a slash.": "Каталог назначения — это каталог, где будут храниться бэкапы. Этот путь задаётся относительно базового каталога назначения. Он должен начинаться с '/', не содержать пробелов и не заканчиваться '/'.",
	"For local backups, the base directory is /backups. So backup files will be stored in:":                                                                                                                      "Для локальных бэкапов базовый каталог — /backups. Поэтому файлы будут храниться в:",
	"For remote backups, the base directory is the bucket root. So backup files will be stored in:":                                                                                                              "Для удалённых бэкапов базовый каталог — корень бакета. Поэтому файлы будут храниться в:",
	"Retention days specifies how many days backup files are kept before automatic deletion. This ensures old backups are removed to save storage space. The retention period is evaluated at execution time.":   "Дни хранения определяют, сколько дней файлы бэкапов хранятся до автоматического удаления. Это помогает удалять старые бэкапы и экономить место. Срок хранения оценивается во время выполнения.",
	"If you set retention days to 0, backups will never be deleted.":                                                                                                                                             "Если установить дни хранения в 0, бэкапы никогда не будут удаляться.",
	"This software uses the battle-tested pg_dump utility to create backups. It creates consistent backups even if the database is being used concurrently.":                                                     "Это приложение использует проверенную утилиту pg_dump для создания бэкапов. Она создаёт согласованные бэкапы даже при одновременной работе с базой данных.",
	"These options are passed to pg_dump. By default, PG Back Web does not pass any options, so backups are full backups.":                                                                                       "Эти параметры передаются в pg_dump. По умолчанию PG Back Web не передаёт параметры, поэтому создаются полные бэкапы.",
	"Learn more in project README":        "Подробнее в README проекта",
	"Learn more in pg_dump documentation": "Подробнее в документации pg_dump",

	// Webhooks
	"Database targets":    "Целевые базы данных",
	"Destination targets": "Целевые хранилища",
	"Backup targets":      "Целевые бэкапы",
	"My webhook":          "Мой вебхук",
	"Event type":          "Тип события",
	"Event types":         "Типы событий",
	"These are the event types that can trigger a webhook.": "Это типы событий, которые могут запускать вебхук.",
	"Database healthy":      "База данных работает",
	"Database unhealthy":    "База данных не работает",
	"Destination healthy":   "Хранилище работает",
	"Destination unhealthy": "Хранилище не работает",
	"Execution success":     "Успешное выполнение",
	"Execution failed":      "Неудачное выполнение",
	"This event is triggered when a database health status changes from unhealthy to healthy.":    "Это событие срабатывает, когда статус базы данных меняется с неработоспособного на работоспособный.",
	"This event is triggered when a database health status changes from healthy to unhealthy.":    "Это событие срабатывает, когда статус базы данных меняется с работоспособного на неработоспособный.",
	"This event is triggered when a destination health status changes from unhealthy to healthy.": "Это событие срабатывает, когда статус хранилища меняется с неработоспособного на работоспособный.",
	"This event is triggered when a destination health status changes from healthy to unhealthy.": "Это событие срабатывает, когда статус хранилища меняется с работоспособного на неработоспособный.",
	"This event is triggered when a backup execution succeeds.":                                   "Это событие срабатывает при успешном выполнении бэкапа.",
	"This event is triggered when a backup execution fails.":                                      "Это событие срабатывает при ошибке выполнения бэкапа.",
	"Activate webhook": "Активировать вебхук",
	"By default it sends a { \"Content-Type\": \"application/json\" } header.": "По умолчанию отправляется заголовок { \"Content-Type\": \"application/json\" }.",
	"By default it sends an empty JSON object {}.":                             "По умолчанию отправляется пустой JSON-объект {}.",

	// Restore
	"Database or connection string is required":                                                         "Требуется база данных или строка подключения",
	"Database and connection string cannot be both set":                                                 "Нельзя указывать одновременно базу данных и строку подключения",
	"Process started, check the restorations page for more details":                                     "Процесс запущен, подробности смотрите на странице восстановлений",
	"Are you sure you want to restore this backup?":                                                     "Вы уверены, что хотите восстановить этот бэкап?",
	"You can restore the backup to an existing database or another database using a connection string.": "Вы можете восстановить бэкап в существующую базу данных или в другую базу по строке подключения.",
	"Select a database": "Выберите базу данных",
	"postgresql://user:password@localhost:5432/mydb": "postgresql://user:password@localhost:5432/mydb",
	"This restoration uses psql v%s":                 "Это восстановление использует psql v%s",
	"Please make sure the target database is compatible with this psql version and double-check that you selected the correct database to restore to.": "Убедитесь, что целевая база данных совместима с этой версией psql, и дважды проверьте, что выбрана правильная база для восстановления.",

	"About PG Back Web": "О PG Back Web",
	"PG Back Web was born in July 2024 from the need for a simple and user-friendly backup solution for self-hosted PostgreSQL databases. After searching extensively for an easy-to-use backup tool and not finding one, I decided to create my own. Its mission is to provide a straightforward web interface that makes managing PostgreSQL backups effortless and efficient.": "PG Back Web появился в июле 2024 года из потребности в простом и удобном решении для резервного копирования self-hosted PostgreSQL. После долгих поисков простого инструмента и не найдя подходящего, я решил создать свой. Его миссия — предоставить понятный веб-интерфейс, который делает управление бэкапами PostgreSQL лёгким и эффективным.",
	"No destinations found":                         "Хранилища не найдены",
	"Wait for the first destination to appear here": "Подождите появления первого хранилища",
	"No databases found":                            "Базы данных не найдены",
	"Wait for the first database to appear here":    "Подождите появления первой базы данных",
}
