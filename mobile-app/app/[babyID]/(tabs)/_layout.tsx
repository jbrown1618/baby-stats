import { Tabs, useLocalSearchParams } from "expo-router";

export default function TabsLayout() {
  const { babyID } = useLocalSearchParams();
  return (
    <Tabs>
      <Tabs.Screen
        name="index"
        options={{ headerShown: false, href: `/${babyID}` }}
      />
      <Tabs.Screen
        name="food"
        options={{ headerShown: false, href: `/${babyID}/food` }}
      />
      <Tabs.Screen
        name="sleep"
        options={{ headerShown: false, href: `/${babyID}/sleep` }}
      />
    </Tabs>
  );
}
