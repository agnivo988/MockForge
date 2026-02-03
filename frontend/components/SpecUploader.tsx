"use client";

export default function SpecUploader() {
  async function upload(e: any) {
    e.preventDefault();

    const form = new FormData(e.target);
    await fetch("http://localhost:8080/api/spec/upload", {
      method: "POST",
      body: form,
    });
    alert("Spec uploaded");
  }

  return (
    <form onSubmit={upload} className="space-y-3">
      <select name="env" className="border p-2">
        <option value="dev">dev</option>
        <option value="staging">staging</option>
        <option value="prod">prod</option>
      </select>

      <input type="file" name="spec" required />
      <button className="bg-black text-white px-4 py-2">
        Upload Spec
      </button>
    </form>
  );
}
