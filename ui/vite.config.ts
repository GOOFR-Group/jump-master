import { defineConfig } from 'vite';
import solid from 'vite-plugin-solid';
import { VitePWA } from 'vite-plugin-pwa';

export default defineConfig({
	plugins: [
		solid(),
		VitePWA({
			registerType: 'autoUpdate',
			strategies: 'generateSW',
			injectRegister: 'inline',
			manifest: {
				name: 'Jump Master',
				start_url: '/',
				display: 'standalone',
				icons: [
					{
						src: 'icon/72.png',
						sizes: '72x72',
						type: 'image/png',
					},
					{
						src: 'icon/128.png',
						sizes: '128x128',
						type: 'image/png',
					},
					{
						src: 'icon/144.png',
						sizes: '144x144',
						type: 'image/png',
					},
					{
						src: 'icon/192.png',
						sizes: '192x192',
						type: 'image/png',
					},
					{
						src: 'icon/512.png',
						sizes: '512x512',
						type: 'image/png',
					},
				],
			},
			workbox: {
				cacheId: 'jump-master',
				globPatterns: ['**/*.{js,css,html,ico,png}'],
			},
		}),
	],
	build: {
		rollupOptions: {
			output: {
				entryFileNames: 'static/js/[name]-[hash].js',
			},
		},
		assetsDir: 'static/css',
	},
});
