Пет-проект "Бот-парсер расписания тренировок в фитнес-клубах".
- MongoDBб 
- gRPC, 
- BurntSushi/toml, 
- sirupsen/logrus, 
- gocolly/colly, 
- go-telegram-bot-api/telegram-bot-api/v5

todo: 
 gRPC между service_bot и service_collect
 1й запрос имен клубов, вывод названий инлайн клавиатурой
 2й запрос по коллбЭку с первой клавиатуры, конкрентого клуба,
 вывод инлайн клавиатуры с днем недели и вывод занятий требуемой даты
 (м.б. кэш в монго)
