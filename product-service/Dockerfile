# Dockerfile for Product Service
# Use a Node.js base image
# FROM node:18-alpine
FROM node:22-alpine

# Set working directory
WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the application files
COPY . .

# Build the NestJS application
RUN npm run build

# Expose the application's port
EXPOSE 50051

# Start the service
CMD ["npm", "run", "start:prod"]
