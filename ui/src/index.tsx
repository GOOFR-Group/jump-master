/* @refresh reload */
import { render } from 'solid-js/web';
import { Router, Route } from '@solidjs/router';

import './index.css';
import Home from './pages/home/home';
import Game from './pages/game/game';
import NotFound from './pages/not-found/not-found';

const root = document.getElementById('root');

render(
	() => (
		<Router>
			<Route path="/" component={Home} />
			<Route path="/index.html" component={Home} />
			<Route path="/play" component={Game} />
			<Route path="*404" component={NotFound} />
		</Router>
	),
	root!,
);
