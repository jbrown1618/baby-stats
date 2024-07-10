import { AddEventButton } from "@/components/AddEventButton";
import { Tabs } from "expo-router";

export default function TabsLayout() {
  return (
    <>
      <Tabs>
        <Tabs.Screen name="index" options={{ headerShown: false }} />
        <Tabs.Screen name="food" options={{ headerShown: false }} />
        <Tabs.Screen name="sleep" options={{ headerShown: false }} />
        <Tabs.Screen name="timeline" options={{ headerShown: false }} />
      </Tabs>
      <AddEventButton />
    </>
  );
}
