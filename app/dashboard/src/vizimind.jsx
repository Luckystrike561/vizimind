export async function listActivity() {
    const res = await fetch(`${import.meta.env.VITE_API_URL}/activities?offset=0&limit=10000`)
    if (!res.ok) {
        console.error('Failed to fetch data');
    }

    return res.json()
}

export async function deleteActivity(item) {
    const res = await fetch(`${import.meta.env.VITE_API_URL}/activity/${item.sku}`, {
        method: 'DELETE'
    })
    if (!res.ok) {
        console.error('Failed to delete data');
    }

    return res.json()
}

export async function upsertActivity(activity) {
    const res = await fetch(`${import.meta.env.VITE_API_URL}/activity`, {
        method: 'POST',
        body: JSON.stringify(activity)
    })
    if (!res.ok) {
        console.error('Failed to create activity');
    }

    return res.json()
}