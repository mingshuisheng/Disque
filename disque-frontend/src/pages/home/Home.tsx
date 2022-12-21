import {Component, For} from "solid-js";
import "./Home.scss"
import {ClassNameUtils} from "../../utils";
import {FlexLayout, Folder} from "../../components";
import {HomeLeft} from "./HomeLeft";
import {HomeTop} from "./HomeTop";
import {fileList} from "../../api";


const {rootClass, className} = ClassNameUtils.create("disque-home");

const Home: Component = () => {

  let handlerClick = async () => {
    let result = await fileList();
    console.log(`name: ${result.name},age: ${result.age}`)
  }

  return (
    <div class={rootClass()}>
      <FlexLayout head={<HomeLeft></HomeLeft>}>
        <FlexLayout direction="column" head={<HomeTop></HomeTop>}>
          <div class={className("file-list")}>
            <For each={new Array<null>(80)}>
              {(e, i) => <Folder onClick={handlerClick} name={`文件夹${i()}`}/>}
            </For>
          </div>
        </FlexLayout>
      </FlexLayout>
    </div>
  )
}

export default Home
