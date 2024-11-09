# Use an official Node.js image as the base
FROM node:18

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json for dependency installation
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Compile TypeScript code
RUN npm run build

# Expose the port specified in the environment variable or default to 3000
EXPOSE ${PORT}

# Set the default command to run the compiled app
CMD ["node", "dist/server/server.js"]
