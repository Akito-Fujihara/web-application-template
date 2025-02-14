'use client'

import {
  Anchor,
  Avatar,
  Button,
  Center,
  Container,
  Group,
  Paper,
  PasswordInput,
  Text,
  TextInput,
} from '@mantine/core'

function Login() {
  return (
    <Center style={{ height: '100vh' }}>
      <Container size={420} my={40}>
        <Avatar
          size={100}
          src="/images/logo.png"
          style={{ margin: '0 auto' }}
        />
        <Text c="dimmed" size="sm" ta="center" mt={5}>
          アカウント未作成の方はこちら{' '}
          <Anchor
            size="sm"
            component="button"
            onClick={() => (window.location.href = '/account/signup')}
          >
            新規アカウント作成
          </Anchor>
        </Text>

        <Paper withBorder shadow="md" p={30} mt={30} radius="md">
          <Text ta="center" fw={700} size="xl">
            ログイン
          </Text>
          <TextInput
            label="メールアドレス"
            placeholder="example@gmail.com"
            required
          />
          <PasswordInput
            label="パスワード"
            placeholder="password"
            required
            mt="md"
          />
          <Group justify="space-between" mt="lg">
            <Anchor component="button" size="sm" style={{ margin: '0 auto' }}>
              パスワードをお忘れですか?
            </Anchor>
          </Group>
          <Button fullWidth mt="xl">
            ログイン
          </Button>
        </Paper>
      </Container>
    </Center>
  )
}

export default Login
