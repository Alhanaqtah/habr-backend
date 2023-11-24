import React from "react";
import { PersonCard } from "../Components/index";
import avatarPlaceholder from "../assets/user.png";

function Following() {
  return (
    <>
      <div className="flex flex-row justify-center my-2">
        <h2 className="text-center text-3xl font-bold text-gray-800">
          Following
        </h2>
      </div>
      <div className="grid grid-flow-row grid-cols-4 gap-y-4 justify-items-center items-center">
        <PersonCard
          userAvatar={avatarPlaceholder}
          userName="Alex"
          userDescription="Lorem"
        />
        <PersonCard
          userAvatar={avatarPlaceholder}
          userName="Alex"
          userDescription="Lorem"
        />
        <PersonCard
          userAvatar={avatarPlaceholder}
          userName="Alex"
          userDescription="Lorem"
        />
        <PersonCard
          userAvatar={avatarPlaceholder}
          userName="Alex"
          userDescription="Lorem"
        />
        <PersonCard
          userAvatar={avatarPlaceholder}
          userName="Alex"
          userDescription="Lorem"
        />
      </div>
    </>
  );
}

export default Following;
