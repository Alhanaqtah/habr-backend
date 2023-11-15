import React from "react";
import { Link } from "react-router-dom";

function User({ userIcon }) {
  return (
    <>
      <div className="w-full h-[90vh] flex flex-col items-center justify-evenly">
        <div className="w-[90%] h-[60%] flex justify-between items-center">
          <div className="bg-gray-200 w-[30%] h-full rounded-xl flex flex-col justify-evenly items-center">
            <img
              src={userIcon}
              alt="user icon"
              className="rounded-[50%] w-[50%]"
            />
            <div className="text-2xl my-2">JoDiFo</div>
          </div>

          <div className="bg-gray-200 w-[65%] h-full rounded-xl"></div>
        </div>

        <div className="w-[90%] h-[25%] flex justify-between">
          <div className="bg-gray-200 w-[32%] h-full rounded-xl">
            <Link to="followers" className="flex w-full h-full">
              <div className="flex flex-col justify-center items-center w-full">
                <div className="text-2xl text-center">Followers</div>
                <hr className="bg-gray-400 w-[60%] h-px my-1" />
                <div className="text-center">0</div>
              </div>
            </Link>
          </div>

          <div className="bg-gray-200 w-[32%] h-full rounded-xl">
            <Link to="following" className="flex w-full h-full">
              <div className="flex flex-col justify-center items-center w-full">
                <div className="text-2xl text-center">Following</div>
                <hr className="bg-gray-400 w-[60%] h-px my-1" />
                <div className="text-center">0</div>
              </div>
            </Link>
          </div>

          <div className="bg-gray-200 w-[32%] h-full rounded-xl">
            <Link to="publications" className="flex w-full h-full">
              <div className="flex flex-col justify-center items-center w-full">
                <div className="text-2xl text-center">Publications</div>
                <hr className="bg-gray-400 w-[60%] h-px my-1" />
                <div className="text-center">0</div>
              </div>
            </Link>
          </div>
        </div>
      </div>
    </>
  );
}

export default User;
