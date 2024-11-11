FROM node:18-bullseye

# Install Bun
RUN curl -fsSL https://bun.sh/install | bash

# Set Bun's path
ENV PATH="/root/.bun/bin:$PATH"

# Set the working directory
WORKDIR /app

# Copy package and lock files
COPY package.json /app/

# Install dependencies
RUN bun install

# Copy the rest of the application code
COPY . /app

# Expose the port the app runs on
EXPOSE 3000

ENV PATH ./node_modules/.bin/:$PATH

# Start the application
CMD ["bun", "run", "dev"]