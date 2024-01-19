import { useState } from "react";
import "./App.css";

function App() {
  const [fileReader] = useState(new FileReader());
  const [file, setFile] = useState(null);

  const handleClick = async () => {
    if (!file) {
      console.error("No file selected");
      return;
    }
    fileReader.readAsArrayBuffer(file);
  };

  const handleFileChange = (event) => {
    const selectedFile = event.target.files[0];
    console.log(selectedFile);
    setFile(selectedFile);
  };

  fileReader.onload = async (e) => {
    const chunkSize = 1000;
    const totalChunks = Math.ceil(e.target.result.byteLength / chunkSize);
    console.log("Read successful");
    const fileName = Math.random() * 1000 + file.name;

    for (let chunkId = 0; chunkId < totalChunks; chunkId++) {
      const start = chunkId * chunkSize;
      const end = Math.min((chunkId + 1) * chunkSize, e.target.result.byteLength);
      const chunk = e.target.result.slice(start, end);

      const response = await fetch("http://localhost:3001/upload", {
        method: "POST",
        headers: {
          "Content-Type": "application/octet-stream",
          "Content-Length": chunk.byteLength.toString(),
          "File-Name": fileName,
        },
        body: chunk,
      });
      console.log(response);
    }

    console.log(e.target.result.byteLength);
  };

  return (
    <>
      <div>
        <h1>File Uploader</h1>
        <input type="file" onChange={handleFileChange} />
        <button onClick={handleClick}>Upload</button>
      </div>
    </>
  );
}

export default App;
