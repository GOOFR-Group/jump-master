/** @type {import('tailwindcss').Config} */
export default {
	content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
	theme: {
		extend: {
			dropShadow: {
				text: '0 1.2px 1.2px rgba(0, 0, 0, 0.8)',
			},
		},
	},
	plugins: [],
};
