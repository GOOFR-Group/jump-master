import { createEffect } from "solid-js";
import { useLocation } from "@solidjs/router";

function NotFound() {
	const location = useLocation();

	createEffect(() => {
    console.log("location:", location);
  });

	return (
		<div>
			Route not found.
		</div>
	);
}

export default NotFound;
