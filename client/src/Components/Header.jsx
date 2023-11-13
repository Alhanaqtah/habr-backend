import React from "react";
import { Link } from "react-router-dom";

function Header({ userIcon }) {
  return (
    <header className="w-full shadow-md sticky top-0 px-8 py-3 bg-white z-10">
      <div className="flex justify-between items-center">
        <div className="flex">
          <nav className="flex gap-[7rem]">
            <div className="text-2xl cursor-pointer m-auto">
              <Link to="articles">Articles</Link>
            </div>
            <div className="text-2xl cursor-pointer m-auto">
              <Link to="feed/articles">Personal articles</Link>
            </div>
            <div className="text-2xl cursor-pointer m-auto">
              <Link to="search">Search</Link>
            </div>
            <div className="text-2xl cursor-pointer m-auto">
              <Link to="users/user">
                <img src={userIcon} alt="profile icon" className="w-[30px]" />
              </Link>
            </div>
          </nav>
        </div>
      </div>
    </header>
  );
}

export default Header;
