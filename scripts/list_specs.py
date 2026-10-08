#!/usr/bin/env python3
import os
from collections import defaultdict

def list_specs():
    specs = defaultdict(list)
    for root, dirs, files in os.walk('specs'):
        for file in files:
            if file.startswith('SPEC') and file.endswith('.md'):
                parts = root.split(os.sep)
                if len(parts) >= 2:
                    category = parts[1]
                    project = '/'.join(parts[2:])
                    if not project:
                        project = '.'
                    specs[category].append(project)

    for category, projects in sorted(specs.items()):
        print(f"📁 {category.capitalize()}:")
        for project in sorted(projects):
            print(f"    📄 {project}")

if __name__ == "__main__":
    list_specs()
