import React from "react";
import { Link } from "react-router-dom";

function Header({ userIcon, searchIcon }) {
  return (
    <header className="w-full shadow-md sticky top-0 py-3 bg-white z-10">
      <div className="flex justify-center items-center w-full">
        <div className="flex w-[75%]">
          <nav className="flex justify-between w-full">
            <div className="flex gap-10">
              <div className="text-2xl cursor-pointer m-auto">
                <Link to="articles">Articles</Link>
              </div>
              <div className="text-2xl cursor-pointer m-auto">
                <Link to="feed/articles">Personal articles</Link>
              </div>
            </div>

            <div className="flex gap-10">
              <div className="text-2xl cursor-pointer m-auto">
                <Link to="search">
                  <img
                    src={searchIcon}
                    alt="search icon"
                    className="w-[30px]"
                  />
                </Link>
              </div>
              <div className="text-2xl cursor-pointer m-auto">
                <Link to="users/user">
                  <img src={userIcon} alt="profile icon" className="w-[30px]" />
                </Link>
              </div>
            </div>
          </nav>
        </div>
      </div>
    </header>
  );
}

export default Header;
