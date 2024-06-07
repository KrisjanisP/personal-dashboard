import os
from datetime import datetime

# Prompt user for the "x" value
x_value = input("Enter the value for x: ")

# Get the current timestamp
timestamp = int(datetime.now().timestamp())

# Define the directory and filenames
migrations_dir = 'migrations'
up_filename = f"{timestamp}_{x_value}.up.sql"
down_filename = f"{timestamp}_{x_value}.down.sql"

# Change to the migrations directory
os.chdir(migrations_dir)

# Create the .up.sql file
with open(up_filename, 'w') as up_file:
    pass  # Create an empty file

# Copy the .up.sql file to create the .down.sql file
os.system(f'cp {up_filename} {down_filename}')

# Change back to the original directory
os.chdir('..')

print(f"Created {up_filename} and {down_filename} in the {migrations_dir} directory.")
