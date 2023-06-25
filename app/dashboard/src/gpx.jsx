export async function listGPX() {
    const res = await fetch('http://127.0.0.1:5000/gpx?offset=0&limit=10000')
    if (!res.ok) {
        console.error('Failed to fetch data');
    }

    return res.json()
}

export async function deleteGPX(item) {
    const res = await fetch(`http://127.0.0.1:5000/gpx/${item.sku}`, {
        method: 'DELETE'
    })
    if (!res.ok) {
        console.error('Failed to delete data');
    }

    return res.json()
}

export async function createGPX(sku, gpx) {
    const res = await fetch(`http://127.0.0.1:5000/gpx`, {
        method: 'POST',
        body: JSON.stringify({
            sku: sku,
            gpx: gpx
        })
    })
    if (!res.ok) {
        console.error('Failed to create data');
    }

    return res.json()
}