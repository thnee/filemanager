export let baseURL = "http://127.0.0.1:4000/api";

export function call({path, method, body, options, headers}) {
	let url = `${baseURL}${path}`;

	method = method || "GET";
	headers = headers || {};
	options = options || {};

	"Content-Type" in headers || (headers["Content-Type"] = "application/json");

	"method" in options || (options.method = method);
	"headers" in options || (options.headers = headers);
	"mode" in options || (options.mode = "cors");
	"credentials" in options || (options.credentials = "include");
	"body" in options || (options.body = JSON.stringify(body));

	return fetch(url, options);
}
