import {Component, For} from 'solid-js';
import {Router, hashIntegration, useRoutes, useNavigate} from "@solidjs/router";
import {routes, setGlobalNavigator} from "./routes";
import {Box, Button, Center, Flex, Stack, useColorMode} from "@hope-ui/solid";

const App: Component = () => {
  const Route = useRoutes(routes)

  const renderRouter = () => {
    setGlobalNavigator(useNavigate())
    return <Route/>
  }

  let {toggleColorMode} = useColorMode();
  return (
    <Router source={hashIntegration()}>
      <Flex minW="$full">
        <Stack as="nav" overflowY="auto" width="200px" height="100vh" bg="$neutral3" direction="column">
          <Center>
            <Stack marginTop="20px" width="80%" direction="column" gap="20px">
              <Button variant="ghost" colorScheme="info">文件</Button>
              <Button variant="ghost" colorScheme="info">相册</Button>
              <Button variant="ghost" colorScheme="info">收藏夹</Button>
              <Button variant="ghost" colorScheme="info">密码箱</Button>
              <Button variant="ghost" colorScheme="info">订阅</Button>
              <Button variant="ghost" colorScheme="info">回收站</Button>
              <Button variant="ghost" colorScheme="info" onClick={toggleColorMode}>换皮肤</Button>
            </Stack>
          </Center>
        </Stack>
        <Flex as="main" flex={1}>
          {renderRouter}
        </Flex>
      </Flex>
    </Router>
  )
};

export default App;
