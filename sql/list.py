import os
import json
import argparse

def list_files_and_contents(directory):
    files_overview = []
    
    for root, dirs, files in os.walk(directory):
        for file in files:
            file_path = os.path.join(root, file)
            try:
                with open(file_path, 'r', encoding='utf-8') as f:
                    content = f.read()
            except Exception as e:
                content = f"Error reading file: {e}"
            
            files_overview.append({
                "filename": file_path,
                "content": content
            })
    
    return files_overview

def save_overview_to_json(overview, output_file):
    with open(output_file, 'w', encoding='utf-8') as f:
        json.dump(overview, f, ensure_ascii=False, indent=4)

def main():
    parser = argparse.ArgumentParser(description="Create an overview of all files in a directory")
    parser.add_argument("directory", type=str, help="The path to the directory")

    args = parser.parse_args()

    overview = list_files_and_contents(args.directory)
    save_overview_to_json(overview, 'files_overview.json')

if __name__ == "__main__":
    main()
