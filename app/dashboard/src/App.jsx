import { createSignal, createResource, Show, For } from "solid-js";
import { upsertActivity, listActivity, deleteActivity } from "./vizimind";

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
  const [activity, setActivity] = createSignal({
    sku: "",
    gpx: "",
    product_id: "",
    names: {
      fr: "",
      en: "",
      de: "",
      sp: ""
    },
    descriptions: {
      fr: "",
      en: "",
      de: "",
      sp: ""
    },
    image: "",
    transports: {
      fr: "",
      en: "",
      de: "",
      sp: ""
    },
    supplier: {
      name: "",
      email: "",
      phone: "",
      address: "",
      city: "",
      zipcode: "",
      country: ""
    },
    extra_meeting_info: {
      fr: "",
      en: "",
      de: "",
      sp: ""
    },
  });
  const [activities] = createResource(listActivity);

  return (
    <div className="mx-auto">
      <form className="grid grid-cols-2 gap-4" action={"/create-gpx"} method='post'>
        <div>
          <label className="block mb-1 font-medium">SKU:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              sku: event.target.value
            });
          }} type="text" value={activity().sku} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter SKU" />
        </div>
        <div>
          <label className="block mb-1 font-medium">GPX:</label>
          <input onChange={(event) => {
            var reader = new FileReader();
            reader.readAsText(event.target.files[0], 'UTF-8');
            reader.onload = function (evt) {
              setActivity({
                ...activity(),
                gpx: evt.target.result
              });
            }
          }} type="file" id="gpx" name="gpx" className="w-full" accept=".xml" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Product ID:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              product_id: event.target.value
            });
          }} type="text" value={activity().product_id} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter Product ID" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity name FR:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              names: {
                ...activity().names,
                fr: event.target.value
              }
            });
          }} type="text" value={activity().names.fr} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the french activity name" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity name EN:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              names: {
                ...activity().names,
                en: event.target.value
              }
            });
          }} type="text" value={activity().names.en} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the english activity name" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity name DE:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              names: {
                ...activity().names,
                de: event.target.value
              }
            });
          }} type="text" value={activity().names.de} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the german activity name" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity name SP:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              names: {
                ...activity().names,
                sp: event.target.value
              }
            });
          }} type="text" value={activity().names.sp} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the spanish activity name" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity description FR:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              descriptions: {
                ...activity().descriptions,
                fr: event.target.value
              }
            });
          }} type="text" value={activity().descriptions.fr} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the french activity description" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity description EN:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              descriptions: {
                ...activity().descriptions,
                en: event.target.value
              }
            });
          }} type="text" value={activity().descriptions.en} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the english activity description" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity description DE:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              descriptions: {
                ...activity().descriptions,
                de: event.target.value
              }
            });
          }} type="text" value={activity().descriptions.de} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the german activity description" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity description SP:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              descriptions: {
                ...activity().descriptions,
                sp: event.target.value
              }
            });
          }} type="text" value={activity().descriptions.sp} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the spanish activity description" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity image:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              image: event.target.value
            });
          }} type="text" value={activity().image} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter an URL" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity transport FR:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              transports: {
                ...activity().transports,
                fr: event.target.value
              }
            });
          }} type="text" value={activity().transports.fr} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the french activity transport" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity transport EN:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              transports: {
                ...activity().transports,
                en: event.target.value
              }
            });
          }} type="text" value={activity().transports.en} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the english activity transport" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity transport DE:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              transports: {
                ...activity().transports,
                de: event.target.value
              }
            });
          }} type="text" value={activity().transports.de} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the german activity transport" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity transport SP:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              transports: {
                ...activity().transports,
                sp: event.target.value
              }
            });
          }} type="text" value={activity().transports.sp} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the spanish activity transport" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity supplier name:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              supplier: {
                ...activity().supplier,
                name: event.target.value
              }
            });
          }} type="text" value={activity().supplier.name} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the supplier name" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity supplier email:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              supplier: {
                ...activity().supplier,
                email: event.target.value
              }
            });
          }} type="text" value={activity().supplier.email} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the supplier email" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity supplier phone:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              supplier: {
                ...activity().supplier,
                phone: event.target.value
              }
            });
          }} type="text" value={activity().supplier.phone} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the supplier phone" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity supplier address:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              supplier: {
                ...activity().supplier,
                address: event.target.value
              }
            });
          }} type="text" value={activity().supplier.address} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the supplier address" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity supplier city:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              supplier: {
                ...activity().supplier,
                city: event.target.value
              }
            });
          }} type="text" value={activity().supplier.city} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the supplier city" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity supplier zipcode:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              supplier: {
                ...activity().supplier,
                zipcode: event.target.value
              }
            });
          }} type="text" value={activity().supplier.zipcode} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the supplier zipcode" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Activity supplier country:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              supplier: {
                ...activity().supplier,
                country: event.target.value
              }
            });
          }} type="text" value={activity().supplier.country} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the supplier country" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Extra meeting info FR:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              extra_meeting_info: {
                ...activity().extra_meeting_info,
                fr: event.target.value
              }
            });
          }} type="text" value={activity().extra_meeting_info.fr} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the french extra meeting info" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Extra meeting info EN:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              extra_meeting_info: {
                ...activity().extra_meeting_info,
                en: event.target.value
              }
            });
          }} type="text" value={activity().extra_meeting_info.en} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the english extra meeting info" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Extra meeting info DE:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              extra_meeting_info: {
                ...activity().extra_meeting_info,
                de: event.target.value
              }
            });
          }} type="text" value={activity().extra_meeting_info.de} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the german extra meeting info" />
        </div>
        <div>
          <label className="block mb-1 font-medium">Extra meeting info SP:</label>
          <input onChange={(event) => {
            setActivity({
              ...activity(),
              extra_meeting_info: {
                ...activity().extra_meeting_info,
                sp: event.target.value
              }
            });
          }} type="text" value={activity().extra_meeting_info.sp} className="w-full px-3 py-2 border border-gray-300 rounded" placeholder="Enter the spanish extra meeting info" />
        </div>
        <div className="col-span-2">
          <button type="submit" className="w-full px-4 py-2 font-medium text-white bg-blue-500 rounded hover:bg-blue-600"
            onClick={
              (e) => {
                e.preventDefault();
                upsertActivity(activity());
                console.log(activity());
              }
            }
          >Submit</button>
        </div>
      </form>
      <Show when={!activities.loading} fallback={<>Loading...</>}>
        <table className="mt-8">
          <thead>
            <tr>
              <th className="px-4 py-2">SKU</th>
              <th className="px-4 py-2">Product ID</th>
              <th className="px-4 py-2">Activity name (FR)</th>
              <th className="px-4 py-2">Supplier name</th>
              <th className="px-4 py-2">Supplier email</th>
              <th className="px-4 py-2">Supplier phone</th>
              <th className="px-4 py-2">Supplier address</th>
              <th className="px-4 py-2">Actions</th>
            </tr>
          </thead>
          <tbody>
            <For each={activities().items} fallback={<div>No activities</div>}>
              {(item) => (
                <tr>
                  <td className="px-4 py-2">{item.sku.toUpperCase()}</td>
                  <td className="px-4 py-2">{item.product_id}</td>
                  <td className="px-4 py-2">{item.names.fr}</td>
                  <td className="px-4 py-2">{item.supplier.name}</td>
                  <td className="px-4 py-2">{item.supplier.email}</td>
                  <td className="px-4 py-2">{item.supplier.phone}</td>
                  <td className="px-4 py-2">{item.supplier.address}</td>
                  <td className="flex space-x-2">
                    <button className="px-4 py-2 text-white bg-blue-500 rounded hover:bg-blue-600" onClick={
                      () => downloadGPX(item)
                    }>Download</button>
                    <button className="px-4 py-2 text-white bg-orange-500 rounded hover:bg-orange-600" onClick={() => {
                      setActivity(item);
                    }
                    }>Edit</button>
                    <button className="px-4 py-2 text-white bg-red-500 rounded hover:bg-red-600" onClick={
                      () => deleteActivity(item)
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
