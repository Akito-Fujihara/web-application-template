import { Button, Center, Container, Stack, Title } from '@mantine/core'

export default function Home() {
  return (
    <Container
      size="sm"
      style={{
        height: '100vh',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
      }}
    >
      <Center>
        <Stack align="center">
          <Title order={1}>Web Application Template</Title>
          <Button size="lg" variant="filled" color="blue" fullWidth>
            ログイン
          </Button>
          <Button size="lg" variant="outline" color="blue" fullWidth>
            新規アカウント作成
          </Button>
        </Stack>
      </Center>
    </Container>
  )
}
