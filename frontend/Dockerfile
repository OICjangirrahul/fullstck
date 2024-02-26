# Use the latest Node.js LTS as a parent image
FROM node:20

# Set the working directory to /app
WORKDIR /app/frontend

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the remaining frontend files to the container
COPY . .

# Expose port 3000
EXPOSE 3000

