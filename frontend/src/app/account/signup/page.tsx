'use client'

import {
  Anchor,
  Avatar,
  Button,
  Center,
  Container,
  Paper,
  PasswordInput,
  Space,
  Text,
  TextInput,
} from '@mantine/core'

function SignUp() {
  return (
    <Center style={{ height: '100vh' }}>
      <Container style={{ width: 360, maxWidth: 360 }} my={40}>
        <Avatar
          size={100}
          src="/images/logo.png"
          style={{ margin: '0 auto' }}
        />
        <Text c="dimmed" size="sm" ta="center" mt={5}>
          アカウント作成済みの方はこちら <br />
          <Anchor
            size="sm"
            component="button"
            onClick={() => (window.location.href = '/account/login')}
          >
            ログイン
          </Anchor>
        </Text>

        <Paper withBorder shadow="md" p={30} mt={30} radius="md">
          <Text ta="center" fw={700} size="xl">
            新規アカウント作成
          </Text>
          <Space h={5} />
          <TextInput label="お名前" placeholder="山田 太郎" required mt="md" />
          <TextInput
            label="メールアドレス"
            placeholder="example@gmail.com"
            required
            mt="md"
          />
          <PasswordInput
            label="パスワード"
            placeholder="password"
            required
            mt="md"
          />
          <PasswordInput
            label="パスワード（確認）"
            placeholder="password"
            required
            mt="md"
          />
          <Button fullWidth mt="xl">
            新規登録
          </Button>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignUp
