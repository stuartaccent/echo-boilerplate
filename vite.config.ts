import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';

export default defineConfig({
	plugins: [svelte()],
	build: {
		outDir: 'pkg/static/public/js',
		emptyOutDir: false,
		rollupOptions: {
			input: 'pkg/ui/lib/index.ts',
			output: {
				entryFileNames: 'ui.js',
				chunkFileNames: '[name].js',
				assetFileNames: '[name].[ext]'
			}
		},
		cssCodeSplit: false,
		minify: true,
	}
});
