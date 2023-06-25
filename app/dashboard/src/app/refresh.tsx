"use client";
import { Item, deleteGPX } from './gpx';

export default function DeleteButton({ item }: { item: Item }) {
    return (
        <button className="px-4 py-2 text-white bg-red-500 rounded hover:bg-red-600" onClick={
            () => deleteGPX(item)
        }>Delete</button>
    )
}
