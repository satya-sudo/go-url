# ---- Build Stage ----
FROM node:18-alpine AS build
WORKDIR /app

# install deps
COPY package*.json ./
RUN npm install

# copy source and build
COPY . .
RUN npm run build

# ---- Serve Stage ----
FROM nginx:alpine
WORKDIR /usr/share/nginx/html

# remove default nginx page
RUN rm -rf ./*

# copy built files from build stage
COPY --from=build /app/dist .

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
