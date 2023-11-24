function PersonCard({ userAvatar, userName, userDescription }) {
  return (
    <>
      <div className="w-[350px] h-[110px] bg-gray-200 rounded-lg flex justify-evenly items-center cursor-pointer hover:bg-gray-300">
        <img
          src={userAvatar}
          alt="course icon"
          className="h-[90px] w-[90px] rounded-sm"
        />
        <div className="flex flex-col justify-start w-[60%] h-[90px]">
          <div className="font-medium text-lg">{userName}</div>
          <div className="font-thin">{userDescription}</div>
        </div>
      </div>
    </>
  );
}

export default PersonCard;
