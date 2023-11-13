import React from "react";
import { PostCard } from "../Components/index";
import hollowKnightImage from "../assets/hollowKnight.png";

function Articles() {
  const text =
    "Lorem ipsum dolor sit amet consectetur adipisicing elit. Vero earum voluptas quaerat, debitis amet quibusdam nobis culpa autem numquam error sint consequuntur quod doloremque minus iusto deleniti quisquam aliquam nemo?";

  const title = "Lorem, ipsum dolor.";

  return (
    <>
      <main className="w-full flex justify-center">
        <div className="w-[75%] flex flex-row justify-center">
          <section className="pt-3 flex flex-col gap-5">
            <PostCard title={title} body={text} imgSrc={hollowKnightImage} />
            <PostCard title={title} body={text} imgSrc={null} />
            <PostCard title={title} body={text} imgSrc={hollowKnightImage} />
          </section>
        </div>
      </main>
    </>
  );
}

export default Articles;
