*.exe
*.exe~
*.dll
*.so
*.dylib
*.test

# Build directories
bin/
build/
dist/

# Compiled object files, caches, etc.
*.o
*.a
*.out

# Go workspace and module caches
go.work
go.work.sum

# Local development environment
.vscode/
.idea/
*.swp

# ===============================
# Logs & Temp Files
# ===============================
*.log
*.tmp
*.pid
*.seed
*.bak
*.old
*.orig

# ===============================
# Environment & Config Files
# ===============================
# Ignore all .env files except example and environment templates
.env
.env.*
!.env.example
!.env.production
!.env.development
!.env.test

# ===============================
# Dependency Manager (if any)
# ===============================
vendor/

# ===============================
# OS-specific files
# ===============================
# macOS
.DS_Store
.AppleDouble
.LSOverride

# Windows
Thumbs.db
ehthumbs.db
Desktop.ini

# Linux
*~

# ===============================
# Other common ignored folders
# ===============================
node_modules/
coverage/
tmp/
cache/
