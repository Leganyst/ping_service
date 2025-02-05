# 1. Используем Node.js для сборки фронта
FROM node:18-alpine AS builder

# 2. Устанавливаем рабочую папку
WORKDIR /app

# 3. Копируем package.json и package-lock.json
COPY package*.json ./

# 4. Устанавливаем зависимости
RUN npm install

# 5. Копируем весь код фронтенда
COPY . .

# 6. Запускаем билд (Vite создаёт статические файлы в dist/)
RUN npm run build

# 7. Используем Nginx для раздачи фронта
FROM nginx:alpine

# 8. Копируем билд фронта в Nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# 9. Копируем конфиг Nginx (если нужно, но в нашем случае фронтенд-сервис служит просто статику)
# Если вы хотите, чтобы фронтенд сам обрабатывал маршруты, можно оставить дефолтный конфиг
# Или можно скопировать кастомный, если требуется:
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# 10. Открываем порт
EXPOSE 80

# 11. Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]
