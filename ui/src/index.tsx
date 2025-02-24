/* @refresh reload */
import { render } from 'solid-js/web';
import { MemoryRouter, Route } from '@solidjs/router';

import './index.css';
import Home from './pages/home/home';
import Play from './pages/play/play';

const root = document.getElementById('root');

render(
	() => (
		<MemoryRouter>
			<Route path="*" component={Home} />
			<Route path="/play" component={Play} />
		</MemoryRouter>
	),
	root!,
);
