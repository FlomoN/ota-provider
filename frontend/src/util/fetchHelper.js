const server = window.location.host;

console.log(window.location.host);

export const get = async (path) => customFetch(path, "GET", undefined);

export const post = async (path, data) => customFetch(path, "POST", data);

export async function customFetch(path, method, data) {
  const res = await fetch("//" + server + path, {
    method: method,
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  return res;
}
