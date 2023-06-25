import { createSignal, createResource, Show, For } from "solid-js";
import { listGPX, deleteGPX, createGPX } from "./gpx";

async function downloadGPX(item) {
  var element = document.createElement('a');
  element.setAttribute('href', 'data:Application/octet-stream,' + encodeURIComponent(item.gpx));
  element.setAttribute('download', `${item.sku}.xml`);

  element.style.display = 'none';
  document.body.appendChild(element);

  element.click();

  document.body.removeChild(element);
}

function App() {
  const [sku, setSku] = createSignal("");
  const [gpx, setGPX] = createSignal("");
  const [listGPXResponse] = createResource(listGPX);

  return (
    <div className="w-96 mx-auto">
      <form className="space-y-4" action={"/create-gpx"} method='post'>
        <div>
          <label className="block mb-1 font-medium">SKU:</label>
          <input onChange={(event) => {
            setSku(event.target.value);
          }} type="text" id="sku" name="sku" className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter SKU" required />
        </div>
        <div>
          <label className="block mb-1 font-medium">GPX:</label>
          <input onChange={(event) => {
            var reader = new FileReader();
            reader.readAsText(event.target.files[0], 'UTF-8');
            reader.onload = function (evt) {
              setGPX(evt.target.result);
            }
          }} type="file" id="gpx" name="gpx" className="w-full" accept=".xml" required />
        </div>
        <div>
          <button type="submit" className="w-full px-4 py-2 font-medium text-white bg-blue-500 rounded hover:bg-blue-600"
            onClick={
              (e) => {
                e.preventDefault();
                createGPX(sku(), gpx());
              }
            }
          >Submit</button>
        </div>
      </form>
      <Show when={!listGPXResponse.loading} fallback={<>Loading...</>}>
        <table className="w-full mt-8">
          <thead>
            <tr>
              <th className="px-4 py-2">SKU</th>
              <th className="px-4 py-2">Actions</th>
            </tr>
          </thead>
          <tbody>
            <For each={listGPXResponse().items} fallback={<div>No GPX</div>}>
              {(item) => (
                <tr>
                  <td className="px-4 py-2">{item.sku.toUpperCase()}</td>
                  <td className="flex space-x-2">
                    <button className="px-4 py-2 text-white bg-blue-500 rounded hover:bg-blue-600" onClick={
                      () => downloadGPX(item)
                    }>Download</button>
                    <button className="px-4 py-2 text-white bg-red-500 rounded hover:bg-red-600" onClick={
                      () => deleteGPX(item)
                    }>Delete</button>
                  </td>
                </tr>
              )}
            </For>
          </tbody>
        </table>
      </Show>
    </div>
  );
}

export default App;
