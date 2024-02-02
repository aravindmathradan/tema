import { browser } from "$app/environment";
import { writable } from "svelte/store";

export const theme = writable((browser && localStorage.theme) || "light");

theme.subscribe((value) => {
	if (browser) return (localStorage.theme = value);
});

export const user = writable();
