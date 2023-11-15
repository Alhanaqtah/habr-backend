import React from "react";

function Footer() {
  return (
    <>
      <div className="bg-gray-800 w-full h-[30vh] mt-4 flex justify-center">
        <div className="w-[75%] h-full py-2 flex flex-col justify-between">
          <div className="w-full h-full">
            <div className="text-2xl text-gray-200 ml-2">Head text 1</div>
            <hr className="w-full h-px bg-gray-200 my-3" />
            <div className="text-gray-200 ml-1 text-lg mb-2">massage 1</div>
            <div className="text-gray-200 ml-1 text-lg mb-2">massage 2</div>
            <div className="text-gray-200 ml-1 text-lg mb-2">massage 3</div>
            <hr className="w-full h-px bg-gray-200 my-3" />
            <div className="text-gray-200 font-light italic">
              Habr-backend 2023
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default Footer;
