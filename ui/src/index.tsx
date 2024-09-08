/* @refresh reload */
import { render } from 'solid-js/web';
import { Router, Route } from '@solidjs/router';

import './index.css';
import Home from './pages/home/home';
import Play from './pages/play/play';

const root = document.getElementById('root');

render(
	() => (
		<Router>
			<Route path="*" component={Home} />
			<Route path="/play" component={Play} />
		</Router>
	),
	root!,
);
