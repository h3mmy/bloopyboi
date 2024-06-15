import { dev } from '$app/environment';
const base = dev ? 'http://localhost:8080' : '';

export function api(
	fetch: Function,
	method: string,
	resource: string,
	data?: Record<string, unknown>
) {
	return fetch(`${base}/${resource}`, {
		method,
		headers: {
			Accepts: 'application/json',
			'Content-Type': 'application/json'
		},
		body: data && JSON.stringify(data)
	});
}
