const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

console.log('🚀 Building project...');

try {
    // Ensure output directory exists
    const distDir = path.join(__dirname, 'views', 'dist');
    if (!fs.existsSync(distDir)) {
        fs.mkdirSync(distDir, { recursive: true });
    }

    // Build with Vite (includes CSS processing via import)
    execSync('npx vite build', {
        stdio: 'inherit'
    });

    console.log('✅ Build completed successfully!');
    console.log('📁 Output files in views/dist/');
} catch (error) {
    console.error('❌ Build failed:', error.message);
    process.exit(1);
}