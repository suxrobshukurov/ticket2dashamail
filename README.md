# 🎟 Ticket Processing Service

Этот сервис **автоматически получает список купленных билетов**, фильтрует email'ы и **отправляет их в Dashamail**.

## 🚀 Функциональность
✅ Получает данные о билетах из API  
✅ Фильтрует email'ы (не отправляет дубли)  
✅ Отправляет email'ы в Dashamail  
✅ Логирует всю работу в файлы  
✅ Работает как фоновый сервис  

---

## 🛠 Установка и запуск

### **1️⃣ Установка Docker**
На сервере выполните:  
```sh
sudo apt update && sudo apt install -y docker.io docker-compose
sudo systemctl enable --now docker
```

### **2️⃣ Клонирование репозитория**
```sh
git clone https://github.com/suxrobshukurov/ticket2dashamail.git
cd ticket-app
```

### **3️⃣ Создание `.env` файла**
Создайте `.env` и добавьте API-ключи:
```ini
DASHAMAIL_API_KEY=your_dashamail_api_key
DASHAMAIL_BASE_ID=your_dashamail_base_id
TICKETS_API_KEY=your_tickets_api_key
LOG_LEVEL=1  # 0 = INFO, 1 = WARNING, 2 = ERROR
```

### **4️⃣ Запуск сервиса**
```sh
docker-compose up -d --build
```
Приложение запустится в фоне.

---

## 📡 API-эндпоинты
🔹 **Проверка состояния сервиса**  
```sh
curl http://localhost:8080/status
```
Ответ: `"Сервис работает"`

🔹 **Принудительный запуск обработки билетов**  
```sh
curl http://localhost:8080/process
```
Ответ: `"Обработка билетов завершена"`

---

## 📜 Логи
Логи сохраняются в `logs/app.log`.  
Для просмотра в реальном времени:  
```sh
tail -f logs/app.log
```

---

## 📌 Остановка сервиса
```sh
docker-compose down
```

---

## 🛠 Разработка и тестирование
🔹 Запуск тестов:
```sh
go test ./...
```
🔹 Локальный запуск без Docker:
```sh
go run cmd/main.go
```

---

## 🤝 Авторы
- **Sukhronjon Shukurov**
- **Ваш [GitHub](https://github.com/suxrobshukurov) / [Telegram](https://t.me/falcon_sukhrob)**
