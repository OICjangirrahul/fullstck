# Use the latest Node.js LTS as a parent image
FROM node:lts

# Set the working directory to /app
WORKDIR /app/backend

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the remaining backend files to the container
COPY . .

# Expose port 8080 (or the port your backend application uses)
EXPOSE 8080

# Command to run the application
CMD ["npm", "start"]
