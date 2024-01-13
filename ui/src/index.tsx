/* @refresh reload */
import { render } from 'solid-js/web';
import { Router, Route } from '@solidjs/router';

import './index.css';
import Home from './pages/home';
import Game from './pages/game';

const root = document.getElementById('root');

render(
	() => (
		<Router>
			<Route path="/" component={Home} />
			<Route path="/game" component={Game} />
		</Router>
	),
	root!,
);
