import {Component, For} from "solid-js";
import "./Home.scss"
import {ClassNameUtils} from "../../utils/ClassNameUtils";
import {Folder} from "../../components";
import {FlexLayout} from "../../components/FlexLayout/FlexLayout";
import {HomeLeft} from "./HomeLeft";
import {HomeTop} from "./HomeTop";


const {rootClass, className} = ClassNameUtils.create("disque-home");

const Home: Component = () => {
  return (
    <div class={rootClass()}>
      <FlexLayout head={<HomeLeft></HomeLeft>}>
        <FlexLayout direction="column" head={<HomeTop></HomeTop>}>
          <div class={className("file-list")}>
            <For each={new Array<null>(80)}>
              {(e, i) => <Folder name={`文件夹${i()}`}/>}
            </For>
          </div>
        </FlexLayout>
      </FlexLayout>
    </div>
  )
}

export default Home
