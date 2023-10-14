


export async function getPost(): Promise<any> {
    const res = await fetch("https://jsonplaceholder.typicode.com/posts", {
        method: 'GET'
    });

    const data = await res.json();

    return data;
}