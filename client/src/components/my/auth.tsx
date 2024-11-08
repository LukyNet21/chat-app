import React from 'react'
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Label } from '../ui/label'
import { Input } from '../ui/input'
import { Button } from '../ui/button'


const AuthComponent = () => {
  return (
    <Tabs defaultValue="login" className="w-[500px] m-auto">
    <TabsList className='mx-autogrid w-full grid-cols-2'>
      <TabsTrigger value="login" className='w-full'>Login</TabsTrigger>
      <TabsTrigger value="register" className='w-full'>Register</TabsTrigger>
    </TabsList>
    <TabsContent value="login">
        <form onSubmit={async (e) => {
          e.preventDefault();
          const formData = new FormData(e.currentTarget);
          const uid = formData.get('uid')?.toString() || '';
          const password = formData.get('password')?.toString() || '';

          try {
            const response = await fetch('http://localhost:8080/api/login', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json'
              },
              body: JSON.stringify({
                userid: uid,
                password: password
              })
            });

            if (!response.ok) {
              throw new Error('Login failed');
            }

            window.location.href = '/'; // Redirect after successful login
          } catch (err) {
            console.error('Login error:', err);
            alert('Login failed');
          }
        }}>
          <Label htmlFor='uid'>User ID</Label>
          <Input id='uid' required></Input>
          <Label htmlFor='password'>Password</Label>
          <Input id='password' type='password' required></Input>
          <Button type='submit' className='mt-2'>Login</Button>
        </form>
    </TabsContent>
    <TabsContent value="register">
       <form onSubmit={async (e) => {
         e.preventDefault();
         const formData = new FormData(e.currentTarget);
         const password = formData.get('password')?.toString() || '';
         const password2 = formData.get('password2')?.toString() || '';
         const publicKey = formData.get('publicKey')?.toString() || '';

         console.log('Form values:', password, password2, publicKey);

         if (password !== password2) {
           alert('Passwords do not match');
           return;
         }
         try {
           const response = await fetch('http://localhost:8080/api/register', {
             method: 'POST',
             headers: {
               'Content-Type': 'application/json'
             },
             body: JSON.stringify({
               password: password,
               public_key: publicKey
             }),
             credentials: 'include'
           });

           if (!response.ok) {
             throw new Error('Registration failed');
           }

           const userId = await response.json();
           console.log('Registered with user ID:', userId);

           window.location.href = '/'; // Redirect after successful registration
         } catch (err) {
           console.error('Registration error:', err);
           alert('Registration failed');
         }
         }}>
         <Label htmlFor='password'>Password</Label>
         <Input id='password' name='password' type='password' required />
         <Label htmlFor='password2'>Password again</Label>
         <Input id='password2' name='password2' type='password' required />
         <Label htmlFor='publicKey'>Public Key</Label>
         <Input id='publicKey' name='publicKey' type='text' required />
         <Button type='submit' className='mt-2'>Register</Button>
       </form>
    </TabsContent>
  </Tabs>
  )
}

export default AuthComponent
