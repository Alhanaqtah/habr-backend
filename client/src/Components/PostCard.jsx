import React from "react";

function PostCard({ title, body, imgSrc }) {
  return (
    <div className="bg-gray-200 w-full px-5 py-3 shadow-lg">
      <div className="text-xl mb-3">{title}</div>
      {/* <img src={imgSrc} className="w-[80%] m-auto" /> */}
      <img src={imgSrc} />
      <div className="mt-3">{body}</div>
    </div>
  );
}

export default PostCard;
